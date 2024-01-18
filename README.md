# Project gogolook-test

- Test Host: https://gogolook-test.fly.dev/
- Web version built on HTMX: https://gogolook-test.fly.dev/web 

implement a restful task API application, including the following endpoints: 

| method | endpoint | param | response | description |
| --- | --- | --- | --- | --- |
| GET | tasks | - | `{tasks: Task[]}` | 取得目前所有task清單 |
| POST | tasks | `{tasks: Task[]}` | `{tasks: Task[]}` | 批次新增task |
| PUT | tasks/{id} | `{name: string, status: int}` | `{task: Task}` | 編輯指定task |
| DELETE | tasks/{id} | - | `{name: string}` | 移除指定task |

*Task:
```
{
    id: "08c7200b-7638-4fe6-abd8-c352b4b1ee9f" // generated uuid
    name: "test",   // task name
    status: 0       // 0 represents an incomplete task, while 1 represents a completed task
}
```


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

generate templ component
```bash
make generate
```

generate templ component - watch mode
```bash
make generate-watch
```
