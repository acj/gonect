include $(GOROOT)/src/Make.inc

PKGDIR=$(GOROOT)/pkg/$(GOOS)_$(GOARCH)

TARG=gonect
CGOFILES=gonect.go
#GOFILES=\
#	main.go

include $(GOROOT)/src/Make.pkg

CLEANFILES+=main $(PKGDIR)/$(TARG).a

main: install main.go
	$(GC) main.go
	$(LD) -o $@ main.$O