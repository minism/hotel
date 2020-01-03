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
    
    def spawn(self, gameId):
        return requests.post(url('/spawn?gameId=%s' % gameId), headers=self._headers)

    def update(self, sid, data):
        return requests.put(url('/servers/%s' % sid), json=data, headers=self._headers)

    def alive(self, sid):
        return requests.put(url('/servers/%s/alive' % sid), headers=self._headers)
