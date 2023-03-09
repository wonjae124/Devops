# 0. Log Dataset
- 이상치 탐지를 위해, Label이 존재하는 데이터 확보
- github loghub에서 Label 탐색
  - 하둡 분산 시스템의 로그를 저장하는 HDFS 선택

    <img src = "https://github.com/wonjae124/Devops/blob/main/image/loghubs_2.png?raw=true" width=800>
  - github logizer의 HDFS_10k, anormal dataset 확보
  - 아래의 Structured Log
  - 주요 column으로 Component, content, Eventid, EventTemplate으로 구성되어 있음
      <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-09%2021-29-01.png?raw=true" width=800>
 
# 1. Log analysis
- HDFS란?
  - 하둡의 파일시스템, 클러스터이다.
  - NameNode - 메타데이터 관리, 데이터의 저장 위치를 표시
  - DataNode   - 실질적인 데이터 저장소
  -  Secondary NameNode - 데이터 노드가 다운됬을 떄를, 대체할 데이터 백업 저장소
  -  HDFS의 파일은 청크로 분할. 분할된 블록이 DataNode에 저장된다.<br/>  
  -  주어진 데이터셋은 anormal(0)의 개수가 굉장히 적으므로 불균형 데이터셋. 따라서, 이상 탐지에서 성능지표 F1 score 표현
    <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-09%2022-28-02.png?raw=true" width=800>
  

# 2. Log parsing
<img src ="https://github.com/logpai/logparser/blob/master/docs/img/example.png?raw=true" width=800>


- Log message는 unstructred log로 ML에 사용불가
이에, logparse 패키지를 통해, structured log로 전처리
- Content에 block_* 이외에 반복적으로 등장하는 문장을 EventTemplate으로 지정하고, 이를 E1, E2 등 번호를 매겨서 표기
<br/><br/>

# 3. Log Sequence to vectorize

-  논문 log anomaly에 의하면, 연속적인 Event가 처음에는 정상이었다가 특정 순간에 anormal로 판정 될 수 있다. 따라서, Event의 맥락을 파악하는 과정이 필요함.<br/>
- logizer 패키지 모듈에서 여러 structed log에 공통된 BlockId를 mapping하는 방법을 참고하여, EventSequence로 나열함
  <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-09%2021-54-59.png?raw=true">

- NLP에 학습하고자, tokenize, padding을 적용한 후 sequence를 벡터화시킴

- 양방향의 맥락 정보를 유지하는 BiLSTM을 적용하는 모델 적용

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

- [github - loghub ](https://github.com/logpai/loghub)
- [github - logizer](https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-09%2021-54-59.png?raw=true)
- [Papers - LogAnomaly: Unsupervised Detection of
Sequential and Quantitative Anomalies in Unstructured Logs](https://nkcs.iops.ai/wp-content/uploads/2019/06/paper-IJCAI19-LogAnomaly.pdf)
