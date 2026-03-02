# go-playground

Just a sandbox for testing and exploring Go language features.

## Troubleshooting

For the `ping` runner you need to either run as root or add CAP_NET_RAW capability
for the binary:

```console
sudo setcap cap_net_raw=+ep ./main
```

Install `setcap` program on OpenSUSE if you don't already have it:

```console
sudo zypper install libcap-progs
```
