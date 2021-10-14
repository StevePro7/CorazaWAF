coraza-ruleset
14/10/2021

https://github.com/jptosso/coraza-ruleset


Reference
https://github.com/jptosso/coraza-ruleset/issues/1

git clone https://github.com/jptosso/coraza-ruleset
cd coraza-ruleset
docker build . -t crs
docker run crs -name crs



Depeendencies
https://github.com/fzipi/go-ftw

FTW
Framework for Testing WAFs



docker run -it  -p 2019:2019/tcp -p 443:443/tcp -p 80:80/tcp crs:latest


VS Code
load project
strip back Dockerfile

right click Dockerfile
Build image
corazaruleset:latest
Executing task:
docker build --pull --rm -f "Dockerfile" -t corazaruleset:latest "."



docker logs --tail 1000 -f b8bcb29c94a881c9c5188970747f95aecf3508cf4535e09c5a5f434b0ac05b6c


docker run --rm -d  -p 2019:2019/tcp -p 443:443/tcp -p 80:80/tcp jptosso/coraza-waf:latest 



Attach shell
docker exec -it 61e03b95c4b3d9bffc8cef36851b7c6c4f2ae7cd09327c1dd0b9bb681ed9dbe6 sh 


OR
just run coraza-waf in Dockerfile only

curl http://localhost:80
If you see this message, Coraza and Caddy are working fine!

BUT
curl http://localhost:80/denied
<empty>
