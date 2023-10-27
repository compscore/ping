# Ping

## Check Format

```yaml
- name:
  release:
    org: compscore
    repo: mysql
    tag: latest
  target:
  weight:
```

## Parameters

| parameter |      path       |   type   | default  | required | description                                     |
| :-------: | :-------------: | :------: | :------: | :------: | :---------------------------------------------- |
|  `name`   |     `.name`     | `string` |   `""`   |  `true`  | `name of check (must be unique)`                |
|   `org`   | `.release.org`  | `string` |   `""`   |  `true`  | `organization that check repository belongs to` |
|  `repo`   | `.release.repo` | `string` |   `""`   |  `true`  | `repository of the check`                       |
|   `tag`   | `.release.tag`  | `string` | `latest` | `false`  | `tagged version of check`                       |
| `target`  |    `.target`    | `string` |   `""`   |  `true`  | `target to ping`                                |
| `weight`  |    `.weight`    |  `int`   |   `0`    |  `true`  | `amount of points a successful check is worth`  |

## Examples

```yaml
- name: host_a-ping
  release:
    org: compscore
    repo: ping
    tag: latest
  target: 10.{{ .Team }}.1.1
  weight: 1
```
