# New Vegas Linux Mod Manager (nvlmm)

When I get time, this'll maybe be a proper mod manager. For now, it's a helper program to let me
install mods reversibly, by mounting the `Data` directory as the lower layer of an
[OverlayFS](https://www.kernel.org/doc/Documentation/filesystems/overlayfs.txt), and creating a
loop device for mods.

## Usage

Since `mount` requires root, the `nvlmm` executable must be run with `sudo`, or have the
`CAP_SYS_ADMIN` capability (sadly, no specific capability exists for managing mounts...) added via:

```
setcap cap_sys_admin+ep /path/to/nvlmm
```

Run `nvlmm help` for usage information.
