# GOERLI NETWORK TEST

This test attempts to access the Rhombus Network Lighthouse contract 100 times for each of the goerli rpc access points

    "http://goerli.blockscout.com",
    "http://goerli.prylabs.net",
    "https://rpc.goerli.mudit.blog",
    "https://rpc.slock.it/goerli"

inside a 100 x loop:
    A connection is established 
    The peekData function is called on the Lighthouse contract at 0x4E10A95f0bb2fEc6ec1c4296a16420a018a5F9Fe
    If no errors were encountered, the score is incremented.
    