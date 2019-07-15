import unittest

import requests

def url(path):
    return 'http://localhost:3000%s' % path

class Client(object):
    def __init__(self):
        self._token = None

    @property
    def _headers(self):
        return {'X-Session-Token': self._token}

    def identify(self):
        r = requests.post(url('/identify'))
        self._token = r.json().get('token')
        assert(self._token)
    
    def listServers(self, gameId):
        return requests.get(url('/servers?gameId=%s' % gameId), headers=self._headers)

    def get(self, sid):
        return requests.get(url('/servers/%s' % sid), headers=self._headers)

    def create(self, data):
        return requests.post(url('/servers'), json=data, headers=self._headers)

    def update(self, sid, data):
        return requests.put(url('/servers/%s' % sid), json=data, headers=self._headers)

    def alive(self, sid):
        return requests.put(url('/servers/%s/alive' % sid), headers=self._headers)


class IntegrationTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.client = Client()
        cls.client.identify()

    def _assertError(self, r, msg):
        self.assertEqual(400, r.status_code)
        self.assertRegexpMatches(r.text, msg)

    def testHealthCheck(self):
        r = requests.get(url('/health'))
        self.assertEqual(200, r.status_code)

    def test404(self):
        r = requests.get(url('/foo'))
        self.assertEqual(404, r.status_code)
        r = requests.get(url('/bar'))
        self.assertEqual(404, r.status_code)
    
    def testMissingField(self):
        r = self.client.create({
            'name': '',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self._assertError(r, r'(?i)fail.*name')
    
    def testInvalidFieldType(self):
        r = self.client.create({
            'name': 'name',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 'string',
        })
        self._assertError(r, r'(?i)fail.*port')

        r = self.client.create({
            'name': 123,
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 'string',
        })
        self._assertError(r, r'(?i)fail.*name')
    
    def testCreateAndGet(self):
        r = self.client.create({
            'name': 'foo',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)
        new_id_1 = r.json().get('id')
        self.assertTrue(new_id_1 is not None)

        r = self.client.create({
            'name': 'bar',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        new_id_2 = r.json().get('id')
        self.assertTrue(new_id_2 is not None)

        r = self.client.get(new_id_1)
        self.assertEqual(200, r.status_code)
        self.assertEquals('foo', r.json().get('name'))

        r = self.client.get(new_id_2)
        self.assertEqual(200, r.status_code)
        self.assertEquals('bar', r.json().get('name'))
    
    def testCreateAndList(self):
        r = self.client.create({
            'name': 'foo',
            'gameId': 'gid2',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)

        r = self.client.create({
            'name': 'bar',
            'gameId': 'gid2',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)

        r = self.client.create({
            'name': 'baz',
            'gameId': 'gid3',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)

        r = self.client.listServers('gid2')
        self.assertEqual(200, r.status_code)

        serverNames = [obj['name'] for obj in r.json()]
        self.assertEqual({'foo', 'bar'}, set(serverNames))
    
    def testCreateAndUpdate(self):
        r = self.client.create({
            'name': 'original',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)
        server_id = r.json().get('id')

        r = self.client.update(server_id, {
            'name': 'new name',
            'port': 1001,
        })
        self.assertEqual(200, r.status_code)
        self.assertEqual('new name', r.json().get('name'))
        self.assertEqual(1001, r.json().get('port'))
        
        r = self.client.get(server_id)
        self.assertEqual(200, r.status_code)
        self.assertEqual('new name', r.json().get('name'))
        self.assertEqual(1001, r.json().get('port'))
    
    def testCreateAndPingAlive(self):
        r = self.client.create({
            'name': 'original',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)
        server_id = r.json().get('id')

        r = self.client.alive(server_id)
        self.assertEqual(200, r.status_code)

        r = self.client.alive(server_id)
        self.assertEqual(200, r.status_code)

    def testCantUpdateOtherServer(self):
        # Test that we cant modify a server not owned by this session.
        r = self.client.create({
            'name': 'original',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)
        server_id_1 = r.json().get('id')

        client2 = Client()
        client2.identify()
        r = client2.create({
            'name': 'server2',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)
        server_id_2 = r.json().get('id')

        # We should be able to get each others servers.
        r = self.client.get(server_id_2)
        self.assertEqual(200, r.status_code)
        r = client2.get(server_id_1)
        self.assertEqual(200, r.status_code)

        # We should not be able to update each others servers.
        r = self.client.update(server_id_2, {
            'name': 'new name',
            'port': 1000,
        })
        self.assertEqual(403, r.status_code)
        r = client2.update(server_id_1, {
            'name': 'new name',
            'port': 1000,
        })
        self.assertEqual(403, r.status_code)






if __name__ == '__main__':
    unittest.main()
