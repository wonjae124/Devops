# 0. Aiops Roadmap
- Artifical Intelgence for IT operation
    - Deployment
        -  1. versining tools(Git/Github, svn) : git branch, merge
        -  2. ML/DL Algorithm : training, hyper parameter tuning
        -  3. CI/CD pipeline : github action, jenkins, Gitlab, grafana(for advanced), Prometheus, circleCI
        -  4. Cloud And Devops Automation : Dockers, k8s, public cloud(Aws, GCP, Azure), kubeflow/Airflow, Ansible
          
    
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
- [Continuous integration with Go and GitHub Actions]([https://youtu.be/MpYB4Qcl570 ](https://www.alexedwards.net/blog/ci-with-go-and-github-actions))
- [Go 패키지 생성에서 버전관리 까지](https://breezymind.com/go-semantic-versioning/)
<br/><br/><br/>
