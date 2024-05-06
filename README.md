## Making test deployments easier

You can deploy your application on a public server and host it on a domain of choice without any hassle of configuring it or using containers. Currently only Go is supported.

### How does it work

By leveraging [systemd](https://systemd.io) and [caddy](https://caddyserver.com). The app will create a systemd service of minimal configuration. After that, caddy (a very powerful web server) will start the web server and setup everything including HTTPS.

| Languages  | Supported |
| ---------- | --------- |
| Go         | ✅        |
| Python     | ❌        |
| Ruby       | ❌        |
| JavaScript | ❌        |

### Prerequisites

The `config.yaml` file needs to be adjusted according to your application.

```yaml
app:
  name: example-app # name of your application
  repo_url: https://github.com/thebigbone/load-balancer # github repo
  directory: load-balancer # directory to clone the repo
  language: go # only go is supported for now
  domain: example.com # domain pointed to your public IP
  proxy: localhost:8090 # address of the service running
  arguments: # arguments required for your application
    - "-port"
    - "8090"
```

**Note:** Point your public IP address to the domain name.

### Installation

- install [go](https://go.dev/install)
- install [caddy](https://caddyserver.com/docs/install)
- `cd cmd`
- run `go build .` or `go run .` for running the application.

### TODO

| Tasks                                | Progress |
| ------------------------------------ | -------- |
| Support different languages          |
| Support different init services      |
| Have an option for docker deployment |
| Monitoring and notification          |
| Deploy on every new commit           |
