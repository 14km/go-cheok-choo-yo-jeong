
# 안녕! 척추요정!

- 이 프로그램은 척추 요정이라는 프로그램입니다.
- 람다를 통하여 구현되도록 되어있습니다.
- 주기적으로  Slack에 알람을 줄 수 있도록 구현 되어있습니다.


---
# 배포 방식

```
# 패키지 빌드 
GOOS=linux go build main.go

# AWS 배포를 위한 zip file 생성.
zip function.zip main
```
