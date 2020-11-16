## TODO list of features

### Current techctl features

root@JDALE-Home:/datto/tech# ls -lash
```
* 4.0K -rw-r--r--  1 root root 1.5K May  2  2016 driveFinder.php
* 4.0K -rw-r--r--  1 root root 1.1K May  2  2016 failedFlags.php
* 4.0K -rw-r--r--  1 root root 1.1K May  2  2016 failedScreenshots.php
* 4.0K -rw-r--r--  1 root root 1.3K May  2  2016 help.php
* 4.0K -rw-r--r--  1 root root 1.6K May  2  2016 ifDestroy.php
* 4.0K -rw-r--r--  1 root root  875 May  2  2016 includedVolumes.php
* 8.0K -rw-r--r--  1 root root 4.8K May  2  2016 info.php
* 4.0K -rw-r--r--  1 root root  455 May  2  2016 memory.php
* 4.0K -rw-r--r--  1 root root 3.1K May  2  2016 screenshotFinder.php
* 4.0K -rw-r--r--  1 root root 2.1K May  2  2016 zfs-arc.php
```

Potential features:
Port check - all DWA/SS ports; iscsi/mercury; Like the device network test, but for an agent.
COW file resize
    * requires pulling agent config from the agent `snapctl agent:request <key> config`
get logs and upload to tmp.datto 
fixing the missing volume keys, if eng is not going to push out a fix soon

