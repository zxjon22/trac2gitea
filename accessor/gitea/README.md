# trac2gitea `accessor.gitea` Package

This provides low-level access to the Gitea application.

The interface `Accessor` expresses all of the operations performed on Gitea by the converter.

The default implementation now uses GORM to allow support for target databases other
than sqlite.

