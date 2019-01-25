package main

import ( 
	"bufio"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"9fans.net/go/plumb"
)


func store(m *plumb.Message) {
	filename := m.LookupAttr("filename")
	log.Println(filename)
	// Ensure path exists and is in a writable directory
	if _, err := os.Stat(path.Dir(filename)); os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(filename), 0755)
		log.Println("here")
		if err != nil {
			return
		}
	}
	// Validate URL and fetch
	u, err := url.ParseRequestURI(string(m.Data))
	if err != nil {
		log.Println(err)
		return
	}
	r, err := http.Get(u.String())
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()
	fd, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer fd.Close()
	r.Write(fd)
}

func main() {
	// So we need to really just http.Get that resource, stash it to the resulting location
	// And create any directory we need; so long as it's in the namespace we should be fine.
	fd, err := os.OpenFile("/mnt/plumb/store", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	io := bufio.NewReader(fd)
	m := &plumb.Message{}
	for {
		err := m.Recv(io)
		if err != nil {
			log.Fatal(err)
		}
		store(m)
	}
}
