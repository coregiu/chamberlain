[Auth]
[Add user]
curl "http://localhost:8080/users"  -H "Accept: application/json" -H "Content-type: application/json" -X POST -d "{\"Username\":\"test11\",\"Password\":\"1234\",\"Role\":\"admin\"}"

[Login]
curl "http://localhost:8080/users/login"  -H "Accept: application/json" -H "Content-type: application/json" -X POST -d "{\"Username\":\"test11\",\"Password\":\"123456\"}"

[Query user]
curl http://localhost:8080/users/user/test11 -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"
curl http://localhost:8080/users/count -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"
curl http://localhost:8080/users -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"

[Modify user]
curl "http://localhost:8080/users"  -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d "{\"Username\":\"test11\",\"Password\":\"5678\",\"Role\":\"admin\"}" -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"

[Logout]
curl "http://localhost:8080/users/logout"  -H "Accept: application/json" -H "Content-type: application/json" -X POST -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"


[Delete user]
curl "http://localhost:8080/users"  -H "Accept: application/json" -H "Content-type: application/json" -X DELETE -d "{\"Username\":\"test11\"}" -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"

[Inputs]
[Login]
curl "http://localhost:8080/users/login"  -H "Accept: application/json" -H "Content-type: application/json" -X POST -d "{\"Username\":\"test\",\"Password\":\"123456\"}"

[Get details]
curl "http://localhost:8080/inputs?limit=10&offset=0&year=2020&month=1" -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"
curl "http://localhost:8080/inputs?limit=10&offset=0&year=2020" -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"
curl "http://localhost:8080/inputs?limit=10&offset=0" -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"
curl "http://localhost:8080/inputs/count?year=2020&month=1" -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"
curl "http://localhost:8080/inputs/count?year=2020" -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"
curl "http://localhost:8080/inputs/count" -H "X-AUTH-TOKEN: f3290562-ec3b-4abb-80eb-d5ca999531ee"

[Token]
f3290562-ec3b-4abb-80eb-d5ca999531ee