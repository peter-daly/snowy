# Snowy

The following was automatically generated via [Betwixt](https://github.com/simonrichardson/betwixt).
Date generated on: 2018-02-01T16:30:35Z
# GET /

+ Request
    + Parameters

            resource_id ('9926d88d-df84-413e-85bf-bf6722de5dfc')

    + Headers

            Accept-Encoding: gzip
            User-Agent: Go-http-client/1.1

+ Response 200
    + Headers

            Content-Length: 176
            Content-Type: application/octet-stream
            X-Duration: 82.518µs
            X-Resourceid: 9926d88d-df84-413e-85bf-bf6722de5dfc

    + Body

            TxY_Xw-aYh1ylWbHTRADfE17uwQH0eLGSYGFWthoHQ2G0ekeABZ5OctmlNLEIqzSCKAHKTlIf2mZ650YpEeEBF2H88Z88idG6ZWvWiU2eVG6ov9s1HHEg_FfuQuts3xYIbbZVSakGpUEaAtOfIt2OhsdSdSVXISGIWMlJT_sc43XqeI=

# GET /revisions/

+ Request
    + Parameters

            resource_id ('9926d88d-df84-413e-85bf-bf6722de5dfc')

    + Headers

            Accept-Encoding: gzip
            User-Agent: Go-http-client/1.1

+ Response 200
    + Headers

            Content-Disposition: attachment; filename=9926d88d-df84-413e-85bf-bf6722de5dfc.zip
            Content-Transfer-Encoding: binary
            Content-Type: application/zip
            X-Duration: 17.787µs
            X-Query-Author-Id: 
            X-Query-Tags: 
            X-Resourceid: 9926d88d-df84-413e-85bf-bf6722de5dfc

    + Body

            PK                   @   d1ef60a4e9151df3a6adddeb805bdc8c3c29c9012dbf3a92c272df9fa3524045  ��PK           PK                   @                 d1ef60a4e9151df3a6adddeb805bdc8c3c29c9012dbf3a92c272df9fa3524045PK      n   s     

# POST /

+ Request
    + Parameters


    + Headers

            Accept-Encoding: gzip
            Content-Length: 176
            Content-Type: application/octet-stream
            User-Agent: Go-http-client/1.1

    + Body

            TxY_Xw-aYh1ylWbHTRADfE17uwQH0eLGSYGFWthoHQ2G0ekeABZ5OctmlNLEIqzSCKAHKTlIf2mZ650YpEeEBF2H88Z88idG6ZWvWiU2eVG6ov9s1HHEg_FfuQuts3xYIbbZVSakGpUEaAtOfIt2OhsdSdSVXISGIWMlJT_sc43XqeI=

+ Response 200
    + Headers

            Content-Type: application/json
            X-Duration: 68.308µs

    + Body

            {
                "address": "d1ef60a4e9151df3a6adddeb805bdc8c3c29c9012dbf3a92c272df9fa3524045",
                "content_type": "application/octet-stream",
                "size": 176
            }

