# 0. Text 분류
- BOW 방식은 문서의 빈도수를 기록함, BOW 감성 사전의 경우 사전을 만든다. 동사, 명사, 형용사에 감성을 라벨링해 사전을 만든다. 단, 맥락을 계산하지 않음
- 머신러닝 방식은 라벨이 있는 데이터셋이 있으므로 별도의 사전을 만들지 않는다.
- 딥러닝 RNN은 앞의 정보를 축적해서 뒤에 정보를 예측하는데 사용한다.

# 1. 워드 임베딩
- 단어의 맥락을 학습하는 기법 : 워드 임베딩
- 워드 임베딩 : 단어를 원핫인코딩(범주->수치) 후에 연속된 값을 갖는 축소된 벡터(밀집 벡터-대부분의 변수가 0이 아닌 연속적인 값)으로 변환 
- 카운트 기반은 문서 단위로 임베딩 
- 딥러닝은 임베딩된 단어의 시퀀스로 문서를 표현, 단어의 순서를 고려해 문맥을 파악한다.
- 워드 임베딩시 문서는 2차원 행렬, 1차원 벡터 리스트로 표현된다. 이미 워드 임베딩 결과가 1차원 벡터이기 떄문임

# 1. BOW 문서 임베딩
- 희소 벡터를 밀집 벡터로 변환하지 않는다.

```
$ sudo apt install postgresql postgresql-contrib 

$ sudo systemctl start postgresql.service

$ go mod init tutorial 

$ sudo -u postgres createuser --interactive

$ sudo -u postgres psql 

$ sudo adduser wonjae

$ sudo -i -u wonjae psql

```


### 출처

- [DigitalOceaon - how-to-install-and-use-postgresql-on-ubuntu-22-04 ](https://www.digitalocean.com/community/tutorials/how-
<br><br><br>
