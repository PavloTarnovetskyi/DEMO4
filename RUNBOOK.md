# -------------------------------------- Demo 4 ------------------------------------------
Table of contents 
- [Terratest](#terratest)
- [SupervisorCTL](#supervisorCTL)
- [SonarQube](#sonarQube)
- [Kubernetes](#kubernetes)

# Terratest

Guide to quik start -> [youtube](https://www.youtube.com/watch?v=GLhtnOdSIh0)

Test infrastructure code with Terratest -> [gruntwork](https://terratest.gruntwork.io/)

Terratest modules -> [gihub](https://github.com/gruntwork-io/terratest/tree/master/modules)

Commands for Go:


  
  ```bash
  # must generate go.mod file
  $ go mod init mod

  # to run Go script with tests
  # -v for verbose
  ~ go test -v <name-of-file>.go
  ```
_________________________________________________________________________________________________
# SupervisorCTL
How To Install and Manage Supervisor -> [digitalocean.com](https://www.digitalocean.com/community/tutorials/how-to-install-and-manage-supervisor-on-ubuntu-and-debian-vps)

Main documetation -> [supervisord.org](http://supervisord.org/index.html#)

Linux: supervisor – управление процессами и сервисами -> [rtfm.co.ua](https://rtfm.co.ua/linux-supervisor-upravlenie-processami-i-servisami/#%D0%92%D0%B5%D0%B1-%D0%B8%D0%BD%D1%82%D0%B5%D1%80%D1%84%D0%B5%D0%B9%D1%81)

## Dockerfile
```
FROM alpine:3.15

COPY --chmod=777 ./install.sh /home/install.sh
RUN ./home/install.sh

COPY --chmod=777 ./postgres_service.sh /usr/bin/postgres_service.sh

COPY ./supervisord.conf /etc/supervisor.d/supervisord.ini
# archive already prepared
COPY ./citizen.war /opt/tomcat/latest/webapps/citizen.war

EXPOSE 9001 8080

CMD ["/usr/bin/supervisord"]
```
## install.sh
```bash
#!/bin/sh

apk update && apk add --no-cache postgresql12 supervisor openjdk11 curl
# Postgres
mkdir /var/lib/postgresql/data && chown postgres:postgres /var/lib/postgresql/data
chmod 0700 /var/lib/postgresql/data && mkdir /run/postgresql && chown postgres:postgres /run/postgresql/
su -l postgres -c 'initdb -D /var/lib/postgresql/data'
sed -i "s/#listen_addresses = 'localhost'/listen_addresses = '127.0.0.1'/g" /var/lib/postgresql/data/postgresql.conf

# Create user and db
db_user=<your_DB_user>
db_pass=<your_DB_user_pass>
db_base=ss_demo_1
su -l postgres -c 'pg_ctl start -D /var/lib/postgresql/data'
su -l postgres -c "psql -c \"CREATE ROLE $db_user WITH LOGIN;\""
su -l postgres -c "psql -c \"CREATE DATABASE $db_base;\""
su -l postgres -c "psql -c \"GRANT ALL PRIVILEGES ON DATABASE $db_base to $db_user;\""
su -l postgres -c "psql -c \"ALTER USER $db_user WITH PASSWORD $db_pass;\""

# Tomcat
VERSION=9.0.62
mkdir /opt/tomcat
wget https://www-eu.apache.org/dist/tomcat/tomcat-9/v${VERSION}/bin/apache-tomcat-${VERSION}.tar.gz -P /tmp
tar -xf /tmp/apache-tomcat-${VERSION}.tar.gz -C /opt/tomcat/
ln -s /opt/tomcat/apache-tomcat-${VERSION} /opt/tomcat/latest
chmod +x /opt/tomcat/latest/bin/*.sh



# Supervisor
supervisord -c /etc/supervisord.conf
supervisor_name=admin
supervisor_pass=some_pass
sed -i "s/;\[inet_http_server\]/[inet_http_server]/g" /etc/supervisord.conf
sed -i "s/;port=127.0.0.1/port=*/g" /etc/supervisord.conf
sed -i "s/;username=user/username=$supervisor_name/g" /etc/supervisord.conf
sed -i "s/;password=123/password=$supervisor_pass/g" /etc/supervisord.conf
mkdir /etc/supervisor.d

```

## postgres_service.sh
```bash
#!/bin/sh

# This script is run by Supervisor to start PostgreSQL in foreground mode

function shutdown()
{
    echo "Shutting down PostgreSQL"
    pkill postgres
}

if [ -d /var/run/postgresql ]; then
    chmod 2775 /var/run/postgresql
else
    install -d -m 2775 -o <your_DB_user> -g <your_DB_user_pass> /var/run/postgresql
fi

# Allow any signal which would kill a process to stop PostgreSQL
trap shutdown HUP INT QUIT ABRT KILL ALRM TERM TSTP

exec su -l <your_DB_user> -c "/usr/libexec/postgresql/postgres -D /var/lib/postgresql/data --config-file=/var/lib/postgresql/data/postgresql.conf"

```

## supervisord.conf
```bash
[supervisord]
nodaemon=true

[program:postgres]
command=/usr/bin/postgres_service.sh
autostart=true
autorestart=true
stderr_logfile=/var/log/postgres_err.log
stdout_logfile=/var/log/postgres_out.log
stopsignal=QUIT


[program:tomcat]
command=/opt/tomcat/latest/bin/catalina.sh run
autorestart=true
startsecs=20
stopsignal=INT
stopasgroup=true
killasgroup=true
stdout_logfile=/var/log/catalina.out
stderr_logfile=/var/log/catalina.out
environment=JAVA_HOME="/usr/lib/jvm/default-jvm",JAVA_BIN="/usr/lib/jvm/default-jvm/bin"

```
## Commands

Main commands for building and managing Geo Citizen Docker image:

- for **ubuntu** based:

  ```bash
  # build Docker image
  $ DOCKER_BUILDKIT=1 docker build . -t supervisored_geocitizen
  # run dockerized Geo Citizen infrastructure
  $ ddocker run --name supervisored -p 9002:9001 -p 8081:8080 -d supervisored_geocitizen:latest
  # interaction with Geo Citizen container
  ~ docker exec -it ??? bash
  ```


# SonarQube
Install SonarQube on Ubuntu 20.04 LTS -> [vultr.com](https://www.vultr.com/docs/install-sonarqube-on-ubuntu-20-04-lts/)

SonarScanner -> [sonarqube.org](https://docs.sonarqube.org/latest/analysis/scan/sonarscanner/)

Official SonarQube repo with packages -> [binaries.sonarqube](https://binaries.sonarsource.com/?prefix=Distribution/sonarqube/)




  

After installation&reboot you an access web interfacec of SonarQube on http://public-ip-of-VM:9000/  
Default credeantials are: *admin, admin*  
Upon first login you have to set new password.

### PostgreSQL configuration for Geocit134

Create user and DBs for Geocit134 tests:

  ```bash
  ~ sudo -i -u postgres psql
  ```

  ```sql
  /* in psql cli */
  # CREATE USER <your_DB_user> WITH PASSWORD '<password>';
  # ALTER USER <your_DB_user> CREATEDB;
  # CREATE DATABASE ss_demo_1;
  # CREATE DATABASE ss_demo_1_test;
  # GRANT ALL PRIVILEGES ON DATABASE ss_demo_1 to <your_DB_user>;
  # GRANT ALL PRIVILEGES ON DATABASE ss_demo_1_test to <your_DB_user>;
  ```

Add *md5* auth method for Geocit134 user:

  ```bash
  ~ sudo nano /etc/postgresql/12/main/pg_hba.conf
  ```

  ```ini
  # in pg_hba.conf file
  local   all             geocitizen                              md5
  ```


Have to add in *pom.xml* in project for SonarQube 
```bash
# projectKey for connecting project with SonarQube
<sonar.projectKey>Geocitizen134</sonar.projectKey>
```
```bash
# dependency for sonar-maven-plugin
        <dependency>
            <groupId>org.sonarsource.scanner.maven</groupId>
            <artifactId>sonar-maven-plugin</artifactId>
            <version>3.9.1.2184</version>
        </dependency>
```
```bash
# surfire and jacoco plugins 
<plugin>
   <groupId>org.apache.maven.plugins</groupId>
   <artifactId>maven-surefire-plugin</artifactId>
   <version>3.0.0-M6</version>
   <configuration>
     <testFailureIgnore>true</testFailureIgnore>
   </configuration>
 </plugin>
 
 <plugin>
    <groupId>org.jacoco</groupId>
    <artifactId>jacoco-maven-plugin</artifactId>
    <version>0.8.7</version>
    <executions>
        <!-- Prepares the property pointing to the JaCoCo runtime agent -->
        <execution>
            <id>prepare-agent</id>
            <goals>
                <goal>prepare-agent</goal>
            </goals>
            <!--Uncomment this in case you are using the maven surefire plugin-->
            <!--<configuration>-->
            <!--   &lt;!&ndash; Sets the name of the property containing the settings for JaCoCo-->
            <!--  runtime agent. &ndash;&gt;-->
            <!--  <propertyName>surefireArgLine</propertyName>-->
            <!--</configuration>-->
        </execution>
        <!-- Ensures that the code coverage report for unit tests is created
            after unit tests have been run. -->
        <execution>
            <id>generate-report</id>
            <phase>verify</phase>
            <goals>
                <goal>report</goal>
            </goals>
        </execution>
    </executions>
</plugin>
```

 We have to uncomment `@Ignore` strings in test files.

```
find src/test/java/com/softserveinc/geocitizen -type f -exec sed -i "s/^\\@Ignore/\\/\\/@Ignore/g" {} +
```
### Pipeline for Jenkins
```
pipeline {
    agent any
       stages {
        stage('SCM') {
            steps {
                git 'https://github.com/PavloTarnovetskyi/Geocit134.git'
            }
        }
        stage('build') {
            steps {
                withMaven(maven:'maven') {
                    sh '''#!/bin/bash
                        find src/test/java/com/softserveinc/geocitizen -type f -exec sed -i "s/^\\@Ignore/\\/\\/@Ignore/g" {} + 
                        mvn clean verify
                       '''
                }
            }
        }
        stage('SonarQube analysis') {
            steps {
                withSonarQubeEnv('SonarQube') {
                    // Optionally use a Maven environment you've configured already
                    withMaven(maven:'maven') {
                        sh 'mvn sonar:sonar'
                    }
                }
            }
        }
        stage("Quality Gate") {
            steps {
                timeout(time: 1, unit: 'HOURS') {
                    // Parameter indicates whether to set pipeline to UNSTABLE if Quality Gate fails
                    // true = set pipeline to UNSTABLE, false = don't
                    waitForQualityGate abortPipeline: true
                }
            }
        }
    }
    post {
        always {
            cleanWs()
        }
    }  
}
```
# Kubernetes
How to install Kubernetes on Ubuntu 20.04 LTS -> [infoit.com.ua](https://infoit.com.ua/linux/kak-ustanovit-kubernetes-na-ubuntu-20-04-lts/)

Creating and managemant - PODS -> [youtube](https://www.youtube.com/watch?v=kGwe8IEDiX4&list=PLg5SS_4L6LYvN1RqaVesof8KAf-02fJSi&index=8)

Creating and managemant - DEPLOYMENTS -> [youtube](https://www.youtube.com/watch?v=l2byGad0Kk4&list=PLg5SS_4L6LYvN1RqaVesof8KAf-02fJSi&index=9)

Creating and managemant - SERVICES -> [youtube](https://www.youtube.com/watch?v=MRNGw08i6S4&list=PLg5SS_4L6LYvN1RqaVesof8KAf-02fJSi&index=10)





## SRE & SLA/SLO/SLI (written by Vladyslav Boreyko)

### Theory

SRE methodology explanation (ru) -> [atlassian](https://www.atlassian.com/ru/incident-management/devops/sre)

Simple explanation of SLA/SLO/SLI philosophy (ru) -> [atlassian](https://www.atlassian.com/ru/incident-management/kpis/sla-vs-slo-vs-sli)

Error budget (ru) -> [atlassian](https://www.atlassian.com/ru/incident-management/kpis/error-budget)

**CONCLUSION**: as a result we have:

- SLA - list of agreements with stackeholder about healthy level of quality

- SLO - list of obejectives/goals that a team must hit to mee SLA

- SLI - metrics used to evaluate SLO

- SRE - team/worker that used methodology of SLO/SLI to observe and analys of project

Main object of all this thing is reliability level number presented in the form of percents. Ideal level for project is 99.99% of reliability. The 100% level is utopia and this number is never used.

- What is this methodology in real work ?
- Just specific set of rules to collect metrics from our project (VMs, instances, servers, DBs ...) and calculating them in specific way to get very simple reliability plot of project (with failures) and with reliability level number (mandatory !).

- Do we just need some metric collector (like Prometheus, Graphite, SensuGo ...) and metric visualiser (Grafana, Kibana, Datadog ...) and special set of metric rules to this tools ?
- Yes. Earlier SRE engineer connect and write rules for this metric in raw form (for metric collector and visualiser directly). But now for kit Prometheus-Grafana we can write SLO/SLI in simple form (yaml format) and then get rules for metrics-alerts-metadata in Prometheus form (also yaml format). And this is *slok/sloth* project ...

### Slok/Sloth

Main site of slok/sloth -> [sloth](https://sloth.dev/)

GitHub of slok/sloth -> [github](https://github.com/slok/sloth)

More info about slok/sloth -> [joyk](https://www.joyk.com/dig/detail/1625546903371171)

Examples of slok/sloth API v1 (latest) usage -> [github](https://github.com/slok/sloth/tree/main/pkg/prometheus/api/v1)

slok/sloth is written on Golang so ... -> [pkg.go](https://pkg.go.dev/github.com/slok/sloth@v0.6.0/pkg/prometheus/api/v1)

**INFO**: in fact slok/sloth is very useful and handy converter of simple SLO spec (from slok/sloth yaml 'syntax') to terrific Prometheus rules (also yaml form).

To install slok/sloth -> [sloth](https://sloth.dev/introduction/install/)

- Docker way - actually you have to pack input .yaml file to new Docker image (based on the origin) and then run this new image like

  ```bash
  # you will get yaml output of Prometheus rules to the stdout
  # if you will use '-o' or '--out' to output rules to file - this file will be generated inside of container obviously  
  ~ docker run -it <my-new-sloth-image> generate -i <path-to-file-in-image>
  ```

- Source way - *make build* uses Docker containers while build slok/sloth - this stuff consume almost 2-3GB of RAM at some point

- k8s way - this case is advisable when all project, metric collector, monitoring system are deployed arleady on k8s cluster

After build (Source way) you get *sloth* folder with *bin* folder with *sloth-linux-amd64* binary file (or similar name, depends of your environment because binary file is builded specialy for host system as you can see in it's name). Move this binary to */usr/bin* (rename as *sloth*) with *+x* mode and enjoy).  
*sloth* has '--help' and '--help-long' manuals of course ...

Simple example to first usage of SLO by slok/sloth -> [sloth](https://sloth.dev/introduction/)

- create *some.yaml* file with 'SLO spec'

- call the command

  ```bash
  # to get rules in file
  ~ sloth generate -i some.yaml -o result.yaml

  # to get rules in stdout
  ~ sloth generate -i some.yaml
  ```

- you get Prometheus rules in *result.yaml* file / stdout !

- move the file with rules to your Prometheus 'main' folder - in most cases */etc/prometheus/* (it contain *prometheus.yml* main config file)

- rename file with rules like *prometheus.rules.yml* - this not mandatory but preffer ...

- add entrypoint for this rule file to Prometheus config file *prometheus.yml*

  ```yaml
  rule_files:
  - 'prometheus.rules.yml'
  ```

  - some info about Prometheus rules -> [softwareadept](https://softwareadept.xyz/2018/01/how-to-write-rules-for-prometheus/)

- restart Prometheus service by *systemctl* (if Prometheus is installled like regular app)

- open Prometheus web interface and check new settings from *prometheus.rules.yml* 

  - 'Rules' - new alert rules
  - 'Status-Rules' - new metric rules

- open Grafana web interface

- import new slok/sloth (actually SLO) dashboards from -> [grafana1](https://grafana.com/grafana/dashboards/14348) and [grafana2](https://grafana.com/grafana/dashboards/14643)
  - of cource use Prometheus as source for this dasboards

- enjoy)

## Price calculating

GCP -> [google](https://cloud.google.com/products/calculator)

AWS -> [amazon](https://calculator.aws/#/addService)

- select *Amazon EC2* -> [amazon](https://calculator.aws/#/createCalculator/EC2)

- select ...

Oracle -> [oracle](https://www.oracle.com/cloud/costestimator.html)

## Appendix

### Maven and Bash

Some command for outputs:

  ```bash
  # print test log to file
  ~  mvn test --log-file mvn_test.log
  
  # grep only main line with [heads]
  ~ cat mvn_test.log | grep --color=never '\[INFO\]\|\[ERROR\]\|\[WARNING\]'

  # output log after 'mvn test' in original coloring
  ~ mvn test --log-file my_temp.log && \
    cat my_temp.log | grep --color=never '\[INFO\]\|\[ERROR\]\|\[WARNING\]' |  
    GREP_COLOR='01;34' grep --color=always 'INFO\|$' | 
    GREP_COLOR='01;31' grep --color=always 'ERROR\|$' |  
    GREP_COLOR='01;93' grep --color=always 'WARNING\|$' && \
    rm my_temp.log
  ```

Workflow to demonstrate:

  ```bash
  # origin state
  ~ ./geo.sh
  ~ ./test.sh

  # full tests
  ~ ./unignore_full.sh
  ~ ./test.sh

  # skip 'bad' tests
  ~ ./geo.sh
  ~ ./unignore_part.sh
  ~ ./test.sh

  # only 'good' test files
  ~ ./delete.sh
  ~ ./test.sh
  ```

Command *grep* can color it's output:

- GREP_COLOR='01;34' - greyish blue
- GREP_COLOR='01;93' - yellow
- GREP_COLOR='01;31' - dull red

### SonarQube

Example of SonarQube API usage:

  ```bash
  # generate new token with 'name' for user/owner of '<token>:'
  # this token can be found in the user profile then
  ~ curl -s -X POST --user <token>: http://<url>:9000/api/user_tokens/generate?name=test | jq

  # get status of last analysis of project with key (key can be found in 'Project Information' in 'Project' page)
  ~ curl -s -X POST --user <token>: http://wlados-sonarqube.ddns.net:9000/api/qualitygates/project_status?projectKey=<key> | jq

  # get configs of some Gate by it's id
  ~ curl -s -X POST --user <token>: http://wlados-sonarqube.ddns.net:9000/api/qualitygates/show?id=<id> | jq
  ```

### k8s

Some usefull getters

  ```bash
  ~ sudo kubectl get nodes
  ~ sudo kubectl get namespaces
  ~ sudo kubectl get pods --all-namespaces
  ```

Some usefull setters

  ```bash
  ~ kubectl config set-context --current --namespace=my-namespace

  ```

Get connection string again

  ```bash
  ~ sudo kubeadm token create --print-join-command
  ```

Change role tag

  ```bash
  ~ kubectl label node <name> node-role.kubernetes.io/worker=<new-tag>
  ```

To delete node

- Find the node with ***kubectl get nodes***. We’ll assume the name of node to be removed is “mynode”, replace that going forward with the actual node name.
- Drain it with ***kubectl drain mynode***
- Delete it with ***kubectl delete node mynode***
- On worker node 
  - clean all posible states and configs automatically ***kubeadm reset***
  - clean CNI configs ***sudo rm -rf /etc/cni/net.d/***
  - delete main configs folder ***$HOME/.kube/config***

Namespace vs context -> [stackoverflow](https://stackoverflow.com/questions/61171487/what-is-the-difference-between-namespaces-and-contexts-in-kubernetes)

*Deployment* + *Service* -> [kubernetes](https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/)

About port types -> [bmc](https://www.bmc.com/blogs/kubernetes-port-targetport-nodeport/)

Service for validation k8s yaml files -> [validkube](https://validkube.com/)

Base k8s yaml file explanation -> [Kubernetes](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/)

- apiVersion -> [matthewpalmer](https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-apiversion-definition-guide.html)

- kind -> [medium](https://chkrishna.medium.com/kubernetes-objects-e0a8b93b5cdc)

- spec -> [kubernetes](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)


apiVer

behavior:
  scaleDown:
    stabilizationWindowSeconds: 300
    policies:
    - type: Percent
      value: 100
      periodSeconds: 15
  scaleUp:
    stabilizationWindowSeconds: 0
    policies:
    - type: Percent
      value: 100
      periodSeconds: 15
    - type: Pods
      value: 4
      periodSeconds: 15
    selectPolicy: Max


apiVersion: apps/v1
kind: Deployment
metadata:
  name: php-apache
spec:
  selector:
    matchLabels:
      run: php-apache
  replicas: 1
  template:
    metadata:
      labels:
        run: php-apache
    spec:
      containers:
      - name: php-apache
        image: k8s.gcr.io/hpa-example
        ports:
        - containerPort: 80
        resources:
          limits:
            cpu: 500m
          requests:
            cpu: 200m    


apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: autoscaler-geocitizen
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: geocitizen
  minReplicas: 2
  maxReplicas: 4
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 75            