# CorazaWAF
Temporary repository for Coraza WAF testing

https://github.com/StevePro7/CorazaWAF
https://coraza.io/docs/seclang/directives

docker pull jptosso/coraza-waf
docker run --name my-waf -p 8080:8080 -v /path/to/config:/coraza:ro -d jptosso/coraza-waf

VS Code
docker run --rm -d  -p 2019:2019/tcp -p 443:443/tcp -p 80:80/tcp jptosso/coraza-waf:latest 


https://golangrepo.com/repo/jptosso-coraza-waf-go-security

https://github.com/jptosso/coraza-ruleset


https://medium.com/@jptosso/implementing-coraza-waf-with-docker-a55a995f055e
https://github.com/jptosso/coraza-sample-wordpress


Alternatives
google
coraza waf alternatives golang

https://www.libhunt.com/r/coraza-waf

https://github-dotcom.gateway.web.tr/topics/waf?l=go&o=asc&s=stars


https://github-dotcom.gateway.web.tr/einyx/openwaf-ingress-controller


WAF tester
https://github-dotcom.gateway.web.tr/jreisinger/waf-tester



Coreruleset
https://github.com/coreruleset/coreruleset


CADDY
https://caddyserver.com/docs/automatic-https

Caddy is the first and only web server to use HTTPS automatically and by default



Including OWASP ModSecurity Core Rule Set
https://www.netnea.com/cms/apache-tutorial-7_including-modsecurity-core-rules/


OWASP		Open Web Application Security Project
CRS		Core Rule Set


crs-setup.conf
base configuration 

example
https://github.com/jptosso/coraza-ruleset/blob/master/crs_setup.conf


Rules
https://www.netnea.com/cms/apache-tutorial-7_including-modsecurity-core-rules
e.g.
REQUEST-910-IP-REPUTATION.conf
D:\Github\coreruleset\coreruleset\rules




coraza-waf-docker
Executing task: docker run --rm -it  -p 2019:2019/tcp -p 443:443/tcp -p 80:80/tcp corazawafdocker:latest 

error
standard_init_linux.go:228: exec user process caused: no such file or directory
The terminal process "bash '-c', 'docker run --rm -it  -p 2019:2019/tcp -p 443:443/tcp -p 80:80/tcp corazawafdocker:latest'" terminated with exit code: 1.


whereas the "official" image worked OK
Executing task: docker run --rm -it  -p 2019:2019/tcp -p 443:443/tcp -p 80:80/tcp jptosso/coraza-waf:latest 


2021/10/14 16:43:36.947 INFO    using provided configuration    {"config_file": "/coraza/Caddyfile", "config_adapter": "caddyfile"}
2021/10/14 16:43:36.948 INFO    admin   admin endpoint started  {"address": "tcp/localhost:2019", "enforce_origin": false, "origins": ["localhost:2019", "[::1]:2019", "127.0.0.1:2019"]}
2021/10/14 16:43:36.949 INFO    tls.cache.maintenance   started background certificate maintenance      {"cache": "0xc0003c2000"}
2021/10/14 16:43:36.949 DEBUG   http    starting server loop    {"address": "[::]:80", "http3": false, "tls": false}
2021/10/14 16:43:36.949 INFO    tls     cleaning storage unit   {"description": "FileStorage:/data/caddy"}
2021/10/14 16:43:36.949 INFO    tls     finished cleaning storage units
2021/10/14 16:43:36.949 INFO    autosaved config (load with --resume flag)      {"file": "/config/caddy/autosave.json"}
2021/10/14 16:43:36.949 INFO    serving initial configuration
2021/10/14 16:43:36.949 INFO    watcher watching config file for changes        {"config_file": "/coraza/Caddyfile"}
2021/10/14 16:43:57.048 INFO    shutting down apps, then terminating    {"signal": "SIGTERM"}
2021/10/14 16:43:57.048 WARN    exiting; byeee!! 👋     {"signal": "SIGTERM"}
2021/10/14 16:43:57.049 INFO    tls.cache.maintenance   stopped background certificate maintenance      {"cache": "0xc0003c2000"}
2021/10/14 16:43:57.050 INFO    admin   stopped previous server {"address": "tcp/localhost:2019"}
2021/10/14 16:43:57.050 INFO    shutdown complete       {"signal": "SIGTERM", "exit_code": 0}
