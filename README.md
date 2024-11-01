# MLine

## Team

### Name

0x22E

### Project

Web editor for reproducible pipelines

### Members

- Ksenia Poliakova
- Oleg Sidorenkov
- Anton Timonin
- Egor Timonin
- Tsvetkova Maria Andreevna

## Requirements

### Summary

Web editor for reproducible pipelines designed for creating, editing, validating specifications with a library of predefined and custom types that follow an ML experiment domain model defined as a semantic network represented in YAML. The editor could be used as an extension in VS Code or as Web App (WPA) served from website.

### Stakeholders

- Developers
- DevOps/SRE-engineers
- Q/A-engineers
- Other tech specialists and managers who interact with CI/CD
- Investors
- Society

### Expected Needs

- Delivary releases
- Rollback releases
- Validate and test releases

### Features

- **Viewing specification** - Users have access to their specifications in the web application and also in VS Code extension. So there are two features that are connected with viewing specifications:
	- Viewing YAML - YAML is a user-friendly data serialization format that is well structured, easy to browse and analyze. In this format it is easy for the human eye to see the details of interest.
	- Syntax Highlighting - this feature allows users to avoid confusion while watching specification file, because it also helps human eye to structure information and make it is easy to read and find some defects in realization.

- **Edit specification** - Project provides users with such important possibility like editing configurations, because this is how different defects may be removed. Some details:
	- Editing YAML - As it was written before, YAML is structured, so for users it is very convenient to edit specifications, that are stored in YAML format.
	- Autocompletion - This option is added to the project to help users to edit specifications, so the system can suggest to continue command or job name, etc, so that users do not have to write full commands.

- **Save specification** - Users need to have an option to save their specification files, not to lose some changes that were made. Systems, having option of saving may have two different features: 
	- Manual saves - Users can choose a particular time for saving by themselves, so their specification files might be saved on the particular stage, that users need.
	- Autosaving - System has special feature that helps users not to lose their changes made in specification files in some unexpected situations, such as the electricity went out. So user doesn't have to save everything manually, everything will be saved automatically.

- **Specification Validation** - It is rather important for system to have validation of specification files, that can save much time for users.
	- Customizable linting - Linter scans user's specification files for areas that require to be removed, so it can save user's time.
	- Static Analysis - This feature helps users to detect some bugs or mismatches in specification files.

- **Code Hints** - includes code highlighting, autocompletion and prompts (e. g. entity name automatic suggestions) and highly customizable linting.

- **Repository Integration** - all the changes can be applied to a GitLab/GitHub repository specification file.

- **Framework Integration** - specification state changes affect the MLOps platform via CI/CD framework integration.

### Constraints

- **Security**
The system must provide secure access to specifications. It will be achived via authentication, authorization and ensuring the security of job runners and other infrastructure components. Moreover, we embrace whitelist approach.

- **Maintainability**
The system should be easy to extend with new features and integrations. It will be achieved with multilayered architectury approach of writing code. Also, every entity will have separate non-intersecting functionality which will make system more flexible.

- **Usability**
The platform must support multiple interfaces: VS Code extension and web UI. It lets user to operate efficiently. The UI must be friendly, so user with low experience of using CI/CD can navigate easily. The will also provide a web-editor with auto-completion hints.

- **Reliability**
  - **Availability**: The service have to provide a high level of SLA (e. g. 99.9%) and be reachable/available as much time as possible.
  - **Recoverability**: The system must be able to restore to a functional state after failures within not longer period than. Regular back ups of configurations, results of pipelines and jobs must be provided, too.

- **Observability**
For maintainers debug logging and tracing must be provided. For users, container logs should be provided

## Architecture

### Draft

![Architecture](diagrams/draft-architecture.drawio.svg)

## Integration

### gateway-api

httpPort: 8080
grpcPort: 8082
databaseConnectionString: "postgres://user:password@gateway-db:5432/specifications?sslmode=disable"

gateway-api ходит по grpc integrator-api

### integrator-api

httPort: 8082
grpcPort: 8083

### validator-api

httPort: 8084
grpcPort: 8085


## LLM Prompts
LLM: ChatGPT o1-mini

```
Создай Go-функции, которые используют библиотеку Squirrel для взаимодействия с базой данных. Функции должны быть гибкими и безопасными, принимая в качестве параметров имя таблицы, условия запроса и необходимые данные
для выполнения операций. Обеспечь следующие требования:

1.  Валидация имени таблицы:
    Имя таблицы должно проверяться по белому списку разрешённых таблиц, чтобы предотвратить SQL-инъекции.
    Если имя таблицы не находится в белом списке, функция должна возвращать ошибку.

2.  Построение SQL-запросов с использованием Squirrel:
    Используй методы Squirrel (Select, Insert, Update, Delete) для построения запросов.
    Все условия должны быть параметризованы с помощью Squirrel (Where, Set и т.д.) для предотвращения SQL-инъекций.

3.  Обработка ошибок:
    Корректно обрабатывай ошибки на каждом этапе: построение запроса, выполнение запроса и получение результатов.
    Возвращай понятные и информативные сообщения об ошибках.

4.  Использование контекста:
    Функции должны принимать context.Context для управления временем выполнения и отменой операций.

5.  Примеры функций:
    Функция SELECT:
    Название: GetRecordsByIDs
    Параметры: ctx context.Context, tableName string, id []int64
    Описание: Получает записи из указанной таблицы по ID, выбирая поля id, status.
    Возвращаемые значения: слайс ссылок на model.Specification
```
    
<img width="933" alt="Screenshot 2024-12-01 at 21 49 27" src="https://github.com/user-attachments/assets/4e0c8449-c8cc-4aa5-ae2f-6b21e45f0375">
<img width="933" alt="Screenshot 2024-12-01 at 21 49 40" src="https://github.com/user-attachments/assets/607f6b1f-2d29-4e88-aaf1-499c5f1394ad">
https://github.com/2Delight/mline/commit/88c529d5e10b095eaff7eebab3c4d4423d2ad147


```
“Создай файл Protocol Buffers (.proto) с синтаксисом proto3, пакетом gateway и опцией go_package равной "gateway-api/pkg/gateway". Импортируй файлы "google/api/annotations.proto" и "google/protobuf/timestamp.proto".

Определи сервис GatewayService с следующими методами:
1.	GetSpecification
	Вход: GetSpecificationRequest
	Выход: Specification
	HTTP маршрут: GET "/gateway/specifications/{id}"

2.	UpdateSpecification
	Вход: UpdateSpecificationRequest
	Выход: UpdateSpecificationResponse
	HTTP маршрут: PUT "/gateway/specifications/{id}"

3.	GetStatus
	Вход: GetStatusRequest
	Выход: Status
	HTTP маршрут: GET "/gateway/specifications/{id}/status"

4.	UpdateStatus
	Вход: UpdateStatusRequest
	Выход: StatusUpdateResponse
	HTTP маршрут: POST "/gateway/specifications/{id}/status"

5.	ValidateSpecification
	Вход: ValidateSpecificationRequest
	Выход: ValidationResult
	HTTP маршрут: POST "/gateway/specifications/{id}/validate"

6.	GetHello
	Вход: GetHelloRequest
	Выход: GetHelloResponse
	HTTP маршрут: GET "/ping"

Определи следующие сообщения:

Specification
int64 id = 1;
string name = 2;
string content = 3; // YAML content
string git_path = 4;
string status = 5;
Timestamp created_at = 6;
Timestamp updated_at = 7;

CommitPushResult
string commit_hash = 1;
bool is_success = 2;

MLDevResult
bool is_success = 1;
repeated string artifacts = 2;

Status
int64 id = 1;
string name = 2; // “committed”, “completed”, etc.

StatusUpdate
string new_status = 1;

StatusUpdateResponse
bool is_success = 1;

ValidationResult
bool is_valid = 1;
repeated Error errors = 2;
repeated Hint hints = 3;

Error
string code = 1;
string message = 2;

Hint
string message = 1;

GetSpecificationRequest
int64 id = 1;

UpdateSpecificationRequest
int64 id = 1;
string specification_content = 2;

UpdateSpecificationResponse
bool is_success = 1;

GetStatusRequest
int64 id = 1;

UpdateStatusRequest
int64 id = 1;
StatusUpdate status_update = 2;

ValidateSpecificationRequest
int64 id = 1;
string specification_content = 2;

GetHelloRequest {}

GetHelloResponse
string pong = 1;
```

<img width="933" alt="Screenshot 2024-12-01 at 22 01 25" src="https://github.com/user-attachments/assets/e1cd5831-87a9-4d65-ab1b-ab9b162337df">
<img width="933" alt="Screenshot 2024-12-01 at 22 01 41" src="https://github.com/user-attachments/assets/e9da5d5f-a65d-4e12-80e2-e87833517e2c">
<img width="933" alt="Screenshot 2024-12-01 at 22 01 56" src="https://github.com/user-attachments/assets/f831acd4-d074-491e-9f91-346939a07726">
https://github.com/2Delight/mline/commit/04bf4d8df00b6a677a9dd74662cfb72ec391371c






