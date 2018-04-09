# Mail Related Resources

## Sendmail

### Installation

**On centOS/RHEL 7**: `sudo yum install sendmail sendmail-cf m4`

Some resources on the web to troubleshoot the installation:

- a [nice post](https://tecadmin.net/install-sendmail-server-on-centos-rhel-server) to go through baby steps.
- a [page that list exit codes](https://fossies.org/dox/sendmail.8.15.2/include_2sm_2sysexits_8h.html)
- about [Hostname problems](https://unix.stackexchange.com/questions/239920/how-to-set-the-fully-qualified-hostname-on-centos-7-0#239950): typically when getting such an error while trying to start sendmail: `Mar 12 20:30:43 hostname sendmail[11018]: My unqualified host name (hostname) unknown; sleeping for retry` and the sendmail service does not start. Insure the hostname is correctly defined and in line with `/etc/hosts` file. Remember that only one line should begin with a given ip address