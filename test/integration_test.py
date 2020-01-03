import unittest

import requests

from common import url, Client

class IntegrationTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.client = Client()
        cls.client.identify()

    def _assertError(self, r, msg):
        self.assertEqual(400, r.status_code)
        self.assertRegexpMatches(r.text, msg)
    
    def _assertOk(self, r):
        self.assertEqual(200, r.status_code, r.text)

    def testHealthCheck(self):
        r = requests.get(url('/health'))
        self._assertOk(r)

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
        self._assertOk(r)
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
        self._assertOk(r)
        self.assertEquals('foo', r.json().get('name'))

        r = self.client.get(new_id_2)
        self._assertOk(r)
        self.assertEquals('bar', r.json().get('name'))
    
    def testCreateAndList(self):
        r = self.client.create({
            'name': 'foo',
            'gameId': 'gid2',
            'host': 'www.google.com',
            'port': 1000,
        })
        self._assertOk(r)

        r = self.client.create({
            'name': 'bar',
            'gameId': 'gid2',
            'host': 'www.google.com',
            'port': 1000,
        })
        self._assertOk(r)

        r = self.client.create({
            'name': 'baz',
            'gameId': 'gid3',
            'host': 'www.google.com',
            'port': 1000,
        })
        self._assertOk(r)

        r = self.client.listServers('gid2')
        self._assertOk(r)

        serverNames = [obj['name'] for obj in r.json().get('servers')]
        self.assertEqual({'foo', 'bar'}, set(serverNames))
    
    def testCreateAndUpdate(self):
        r = self.client.create({
            'name': 'original',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self._assertOk(r)
        server_id = r.json().get('id')

        r = self.client.update(server_id, {
            'name': 'new name',
            'port': 1001,
        })
        self._assertOk(r)
        self.assertEqual('new name', r.json().get('name'))
        self.assertEqual(1001, r.json().get('port'))
        
        r = self.client.get(server_id)
        self._assertOk(r)
        self.assertEqual('new name', r.json().get('name'))
        self.assertEqual(1001, r.json().get('port'))

    def testCreateAndUpdateSubsetFields(self):
        r = self.client.create({
            'name': 'original',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self._assertOk(r)
        server_id = r.json().get('id')

        r = self.client.update(server_id, {
            'port': 1001,
        })
        self._assertOk(r)
        self.assertEqual('original', r.json().get('name'))
        self.assertEqual(1001, r.json().get('port'))

        r = self.client.update(server_id, {
            'name': 'new name',
        })
        self._assertOk(r)
        self.assertEqual('new name', r.json().get('name'))
        self.assertEqual(1001, r.json().get('port'))

        # An "Empty" update request should suffice to just ping the server being alive.
        r = self.client.update(server_id, {})
        self.assertEqual(200, r.status_code)
    
    def testCantUpdateOtherServer(self):
        # Test that we cant modify a server not owned by this session.
        r = self.client.create({
            'name': 'original',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self._assertOk(r)
        server_id_1 = r.json().get('id')

        client2 = Client()
        client2.identify()
        r = client2.create({
            'name': 'server2',
            'gameId': 'gid',
            'host': 'www.google.com',
            'port': 1000,
        })
        self._assertOk(r)
        server_id_2 = r.json().get('id')

        # We should be able to get each others servers.
        r = self.client.get(server_id_2)
        self._assertOk(r)
        r = client2.get(server_id_1)
        self._assertOk(r)

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
    
    def testCantSpawnByDefault(self):
        r = self.client.spawn('gid')
        self._assertError(r, 'spawn')


if __name__ == '__main__':
    unittest.main()
