import unittest

import requests

def url(path):
    return 'http://localhost:3000%s' % path

def listServers(gameId):
    return requests.get(url('/servers?gameId=%s' % gameId))

def get(sid):
    return requests.get(url('/servers/%s' % sid))

def create(data):
    return requests.post(url('/servers'), json=data)

def update(sid, data):
    return requests.put(url('/servers/%s' % sid), json=data)

def alive(sid):
    return requests.put(url('/servers/%s/alive' % sid))


class IntegrationTest(unittest.TestCase):
    def setUp(self):
        pass

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
        r = create({
            'name': '',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self._assertError(r, r'(?i)fail.*name')
    
    def testInvalidFieldType(self):
        r = create({
            'name': 'name',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 'string',
        })
        self._assertError(r, r'(?i)fail.*port')

        r = create({
            'name': 123,
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 'string',
        })
        self._assertError(r, r'(?i)fail.*name')
    
    def testCreateAndGet(self):
        r = create({
            'name': 'foo',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)
        new_id_1 = r.json().get('id')
        self.assertTrue(new_id_1 is not None)

        r = create({
            'name': 'bar',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        new_id_2 = r.json().get('id')
        self.assertTrue(new_id_2 is not None)

        r = get(new_id_1)
        self.assertEqual(200, r.status_code)
        self.assertEquals('foo', r.json().get('name'))

        r = get(new_id_2)
        self.assertEqual(200, r.status_code)
        self.assertEquals('bar', r.json().get('name'))
    
    def testCreateAndList(self):
        r = create({
            'name': 'foo',
            'gameId': 'gid2',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)

        r = create({
            'name': 'bar',
            'gameId': 'gid2',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)

        r = create({
            'name': 'baz',
            'gameId': 'gid3',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)

        r = listServers('gid2')
        self.assertEqual(200, r.status_code)

        serverNames = [obj['name'] for obj in r.json()]
        self.assertEqual({'foo', 'bar'}, set(serverNames))
    
    def testCreateAndUpdate(self):
        r = create({
            'name': 'original',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)
        server_id = r.json().get('id')

        r = update(server_id, {
            'name': 'new name',
            'port': 1001,
        })
        self.assertEqual(200, r.status_code)
        self.assertEqual('new name', r.json().get('name'))
        self.assertEqual(1001, r.json().get('port'))
        
        r = get(server_id)
        self.assertEqual(200, r.status_code)
        self.assertEqual('new name', r.json().get('name'))
        self.assertEqual(1001, r.json().get('port'))
    
    def testCreateAndPingAlive(self):
        r = create({
            'name': 'original',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)
        server_id = r.json().get('id')

        r = alive(server_id)
        self.assertEqual(200, r.status_code)

        r = alive(server_id)
        self.assertEqual(200, r.status_code)




if __name__ == '__main__':
    unittest.main()
