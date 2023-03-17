# 0. Aiops Roadmap
- Artifical Intelgence for IT operation
    - Deployment
        -  1. versining tools(Git/Github, svn) : git branch, merge
        -  2. ML/DL Algorithm : training, hyper parameter tuning
        -  3. CI/CD pipeline : github action, jenkins, Gitlab, grafana(for advanced), Prometheus, circleCI
        -  4. Cloud And Devops Automation : Dockers, k8s, public cloud(Aws, GCP, Azure), kubeflow/Airflow, Ansible
- Aiops 엔지니어의 역할
  - IT 운영의 자동화를 위해 AI를 도입하는 엔지니어입니다.
  - 비즈니스에 영향을 미치기 전에, 사전에 이슈를 발견하고 조치를 취하는 운영시스템을 만드는 역할을 수행합니다

- Aiops 엔지니어 수행 업무
    IT인프라를 모니터링하고 분석하는 업무 수행
    - 1. 데이터 수집 및 분석 : IT시스템의 로그, 이벤트 및 성능 데이터를 수집하고 분석하여 이상 현상 및 장애를 탐지하고, 시스템 성능 향상을 위한 정보를 도출합니다.
      - 대상(개인, 단체, 국가)
      - 문제(사내 전산 - 판매보조, 모니터링 및 관리업무, 현장업무) 
      - 해결방법(프로그램)
      - 해결방법의 기능
      - 저장소
    - 2. 자동화된 운영 : IT시스템의 운영을 자동화합니다. 자동화 된 알림 및 대응 기능을 사용하여, 장애에 대한 대응 시간을 단축하고, 비즈니스의 연속성을 보장합니다.
    - 3. 머신러닝 기술 적용: ML/DL 기술을 사용하여, IT시스템의 이상 현상과 장애를 탐지하고 예측합니다. 이를 통해, 사전 예방 및 조취를 취하여 시스템 성능 향상시킵니다
    - 4. 데이터 시각화 및 보고: 데이터를 시각화하고, 대시보드를 제작하여 IT시스템의 상태를 모니터링하고, 분석 결과를 보고합니다.
    - 5. 팀과의 협업 : IT운영팀과 개발팀과 긴밀하게 협력하여, IT시스템의 성능개선 및 장애 대응에 대한 전략을 수집하고, 실행합니다.

-
    
# 1. Github Action-CI/CD
- go언어 파일의 테스트 진행
- .githubs/workflows는 main.go가 있는 곳에 설치해야지 CI 정상 진행됨
- go mod init github.com/<user name>/<repository name> ex) go mod init github.com/wonjae124/devops-blog
- `import "testing"` 테스트 함수를 임의 생성 후, `*testing`으로 매개인자를 받아온다.
    - `main.go`
    ```golang
          package main

          import "fmt"

          func main() {
              msg := sayHello("Alice")
              fmt.Println(msg)
          }

          func sayHello(name string) string {
              return fmt.Sprintf("Hi %s", name)
          }
    
    ```
    - `main_test.go`
        ```golang
            package main

            import "testing"

            func Test_sayHello(t *testing.T) {
                name := "Bob"
                want := "Hello Bob"

                if got := sayHello(name); got != want {
                    t.Errorf("hello() = %q, want %q", got, want)
                }
            }

        ```
     - `go test .`      
        - <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-17%2018-27-38.png" width=1200>

# 느낀점
- 다음부터 협업을 위해서 테스트 코드를 만들어 CI/CD를 관리 해야겠다.
- git으로 go 버젼 관리, branch merge 등 실습 예정
- 

#### 출처
- [Continuous integration with Go and GitHub Actions](https://www.alexedwards.net/blog/ci-with-go-and-github-actions))
- [Go 패키지 생성에서 버전관리 까지](https://breezymind.com/go-semantic-versioning/)
<br/><br/><br/>
- [Youtube - IT](https://youtu.be/xD7BQOyHYjo)
- [Naver blog - wonjae124](https://blog.naver.com/wonjae124/223047931490)
