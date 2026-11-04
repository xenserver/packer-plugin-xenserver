## Examples

Examples of building a Windows image with the XenServer builders. For every option, see the
[xenserver-vm-iso](../docs/builders/xenserver-vm-iso.mdx) and
[xenserver-vm-clone](../docs/builders/xenserver-vm-clone.mdx) documentation.

| Directory | Builder | What it does |
| --- | --- | --- |
| `windows-iso/` | `xenserver-vm-iso` | Installs Windows 11 from an ISO on an ISO SR, then provisions it over WinRM |
| `windows-clone-winrm/` | `xenserver-vm-clone` | Clones an existing Windows template, then provisions it over WinRM |

Keep each example in its own directory: Packer loads every `*.pkr.hcl` in a directory as
one configuration.

### Running an example

Run packer from the plugin's root directory (the one with `go.mod`); the `cd_files`
paths in the templates are relative to it.

1. Export the variables the example needs:

   ```
   PKR_VAR_xenserver_host       # the pool coordinator (master)
   PKR_VAR_xenserver_password
   PKR_VAR_winrm_password
   PKR_VAR_ca_cert_path         # the pool's certificate, see step 2
   ```

   The ISO example also needs `PKR_VAR_iso_name` and `PKR_VAR_iso_sr`: the install ISO and
   the ISO SR it is on. `PKR_VAR_vm_sr` optionally names a writable SR for the VM's disks;
   on a pool of more than one host, `Local storage` matches every host's local SR and
   cannot be used. If your pool names the templates differently, set
   `PKR_VAR_clone_template`.

2. Write out the pool's certificate, and point `PKR_VAR_ca_cert_path` (or
   `XENSERVER_CA_CERT_PATH`) at it:

   ```
   xe host-get-server-certificate host=<host-uuid> > pool-ca.pem
   ```

3. Install the plugin and build:

   ```
   packer init  examples/windows-iso/windows-iso-winrm.pkr.hcl
   packer build examples/windows-iso/windows-iso-winrm.pkr.hcl
   ```

Each example also ships a `*.pkrvars.hcl.example`. Copy it to a name ending in
`.auto.pkrvars.hcl`, which Packer loads automatically, or pass it with `-var-file=<path>`.
Keep passwords out of it.

### Windows

- **WinRM**: both examples connect over HTTPS (`winrm_use_ssl = true`) and verify the
  guest's certificate. The guest's 5986 listener must present a certificate that the
  machine running Packer trusts, issued for the address Packer connects to.
- **XenServer VM Tools**: the build learns the VM's address from the VM Tools, so they must
  be installed in the guest. In `windows-iso/` the answer file installs the MSI you put on
  the `cd_files` CD; in `windows-clone-winrm/` the source template must already have them.
  Without them the build fails after `install_timeout` (200 minutes) with
  `No IP address for the VM`.
- **Answer file**: `windows-iso/` needs an `Autounattend.xml` and the VM Tools MSI on a CD
  via `cd_files`. You supply both; the commented-out `cd_files` line in the template is an
  example path. The answer file must install the VM Tools and create the WinRM listener
  above; see Microsoft's
  [Unattended Windows Setup Reference](https://learn.microsoft.com/en-us/windows-hardware/customize/desktop/unattend/).
- **Windows 11** needs UEFI and a vTPM (`firmware = "uefi"`, `vtpm = true`). On the clone
  builder, `vtpm = false` removes an inherited vTPM and everything sealed to it, such as a
  BitLocker key, and cannot be undone.
