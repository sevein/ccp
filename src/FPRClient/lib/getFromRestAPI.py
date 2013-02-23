#!/usr/bin/python -OO

# This file is part of Archivematica.
#
# Copyright 2010-2013 Artefactual Systems Inc. <http://artefactual.com>
#
# Archivematica is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Archivematica is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Archivematica.  If not, see <http://www.gnu.org/licenses/>.

# @package Archivematica
# @subpackage FPRClient
# @author Joseph Perry <joseph@artefactual.com>
from optparse import OptionParser
from StringIO import StringIO    
import pycurl
import json
import sys
sys.path.append("/usr/lib/archivematica/archivematicaCommon")
from databaseFunctions import insertIntoEvents

def getFromRestAPI(url, postFields, verbose=False):
    USER = "demo"
    PASS = "demo"
    storage = StringIO()
    c = pycurl.Curl()
    c.setopt(c.URL, url)
    #c.setopt(c.POSTFIELDS, postFields)
    if verbose:
        c.setopt(c.VERBOSE, True)
    #c.setopt(pycurl.USERPWD, "%s:%s" % (USER,PASS))
    #c.setopt(pycurl.HTTPAUTH, pycurl.HTTPAUTH_BASIC)
    
    #c.setopt(pycurl.COOKIEJAR, 'cookies.txt')
    #c.setopt(pycurl.COOKIEFILE, 'cookies.txt')
    #c.setopt(pycurl.CAPATH,'FILE.pem')
    #c.setopt(pycurl.SSLCERTPASSWD,'foobar')
    #c.setopt(pycurl.SSLCERTTYPE, 'PEM')
    #c.setopt(pycurl.SSLCERT,'FILE1.pem')
    #c.setopt(pycurl.SSL_VERIFYPEER,0)#--insecure
    
    c.setopt(c.WRITEFUNCTION, storage.write)
    c.perform()
    c.close()
    
    #todo
    if not 200:
        throw()
    
    content = storage.getvalue()
    ret = json.loads(content)
    for x in ret["objects"]:
        print type(x), x

if __name__ == '__main__':
    parser = OptionParser()
    parser.add_option("-u",  "--url",          action="store", dest="url", default="http://fprserver/api/fpr/v1/file_id/")
    parser.add_option("-p",  "--postFields",        action="store", dest="postFields", default="format=json&order_by=lastmodified&lastmodified__gte=2012-10-10T10:00:00")
    parser.add_option("-v",  "--verbose",        action="store_true", dest="verbose", default=False)


    (opts, args) = parser.parse_args()
    getFromRestAPI(opts.url, opts.postFields, verbose=opts.verbose)