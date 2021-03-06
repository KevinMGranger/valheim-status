Name: valwho
Version: 1.0.0
Release: 2
Summary: valheim server presence checker and website
License: MIT
URL: https://github.com/kevinmgranger/valheim-status
BuildRequires: golang systemd-devel
Requires: bash systemd

%description
Tools to parse and determine who's currently in a running valheim server,
and a dynamic web service to list the same info.

%define _specdir %{getenv:PWD}
%define _builddir %{getenv:PWD}

%build
make %{?_smp_mflags}

%install
%make_install

%files
/usr/libexec/valwho/invocation    
/usr/libexec/valwho/logs
/usr/libexec/valwho/who
/usr/libexec/valwho/parse
/usr/libexec/valwho/web
/usr/bin/valwho