# Project gogolook-test

implement a restful task API application, including the following endpoints: 

| method | endpoint | param | response | description |
| --- | --- | --- | --- | --- |
| GET | tasks | - | `{tasks: Task[]}` | 取得目前所有task清單 |
| POST | tasks | `{tasks: Task[]}` | `{isSuccess: boolean, idList: string[]}` | 批次新增task |
| PUT | tasks/{id} | `{name: string, status: int}` | `{isSuccess: boolean, id: string}` | 編輯指定task |
| DELETE | tasks/{id} | - | `{isSuccess: boolean, name: string}` | 移除指定task |

*Task:
```
{
    name: "test",   // task name
    status: 0       // 0 represents an incomplete task, while 1 represents a completed task
}
```

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```