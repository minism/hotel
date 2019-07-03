import unittest

import requests

def url(path):
    return 'http://localhost:3000%s' % path

def create(data):
    return requests.post(url('/servers'), json=data)

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
            'host': 'www.google.com',
            'port': 1000,
        })
        self._assertError(r, r'(?i)fail.*name')
    
    def testInvalidFieldType(self):
        r = create({
            'name': 'name',
            'host': 'www.google.com',
            'port': 'string',
        })
        self._assertError(r, r'(?i)fail.*port')
    
    def testCreateAndGet(self):
        r = create({
            'name': 'foo',
            'host': 'www.google.com',
            'port': 1000,
        })
        self.assertEqual(200, r.status_code)
        new_id_1 = r.json().get('id')
        self.assertTrue(new_id_1)

        r = create({
            'name': 'bar',
            'host': 'www.google.com',
            'port': 1000,
        })
        new_id_2 = r.json().get('id')
        self.assertTrue(new_id_2)

        r = requests.get(url('/servers/%s' % new_id_1))
        self.assertEqual(200, r.status_code)
        self.assertEquals('foo', r.json().get('name'))

        r = requests.get(url('/servers/%s' % new_id_2))
        self.assertEqual(200, r.status_code)
        self.assertEquals('bar', r.json().get('name'))

if __name__ == '__main__':
    unittest.main()
