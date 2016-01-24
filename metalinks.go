package metalinks

import (
	"io/ioutil"
	"encoding/xml"
	"net/http"
)

type Metalink struct {
	XMLName	xml.Name	`xml:"metalink"`
	Text	string		`xml:",chardata"`
	Files	Files
}

type Files struct {
	XMLName		xml.Name	`xml:"files"`
	Text		string		`xml:",chardata"`
	File		[]MetaFile		`xml:"file"`
}

type MetaFile struct {
	XMLName		xml.Name	`xml:"file"`
	Name		string		`xml:"name,attr"`
	Text		string		`xml:",chardata"`
	Size		Size		`xml:"size"`
	Resources	[]Resources	`xml:"resources"`
	Verification	FileVerification	`xml:"verification"`
}

type FileVerification struct {
	XMLName		xml.Name	`xml:"verification"`
	Hashes		[]Hash		`xml:"hash"`
}

type Hash struct {
	XMLName		xml.Name	`xml:"hash"`
	Type		string		`xml:"type,attr"`
	Text		string		`xml:",chardata"`
}

type Size struct {
	XMLName		xml.Name	`xml:"size"`
	Text		string		`xml:",chardata"`
}

type Resources struct {
	XMLName		xml.Name	`xml:"resources"`
	Urls		[]Url		`xml:"url"`
}

type Url struct {
	XMLName		xml.Name	`xml:"url"`
	Type		string		`xml:"type,attr"`
	Protocol	string		`xml:"protocol,attr"`
	Location	string		`xml:"location,attr"`
	Preference	string		`xml:"preference,attr"`
	Link		string		`xml:",chardata"`
}

// Reads a local file and returns a metalink struct
func ParseFile(filename string)(Metalink, error){
	metafile := Metalink{}
	xmlContent, _ := ioutil.ReadFile(filename)

	err := xml.Unmarshal(xmlContent, &metafile)
	if err != nil {
		return metafile, err
	}
	return metafile, nil
}

// Reads a metalink xml string and returns a metalink struct
func ParseString(content string)(Metalink, error){
	metafile := Metalink{}

	err := xml.Unmarshal([]byte(content), &metafile)
	if err != nil {
		return metafile, err
	}
	return metafile, nil
}

// Reads a metalink xml bytes slice and returns a metalink struct
func ParseBytes(content []byte)(Metalink, error){
	metafile := Metalink{}

	err := xml.Unmarshal(content, &metafile)
	if err != nil {
		return metafile, err
	}
	return metafile, nil
}

// Downloads a metalink xml file by a http\https url and returns a metalink struct
func ParseFileFromUrl(link string)(Metalink, error){
	metafile := Metalink{}

	resp, err := http.Get(link)
	if err != nil {
		return metafile, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return metafile, err
	}

	err = xml.Unmarshal([]byte(content), &metafile)
	if err != nil {
		return metafile, err
	}
	return metafile, nil
}

/*
	//fmt.Println(metafile)
	fmt.Println(metafile.Files.File[0].Name)

	fmt.Println(metafile.Files.File[0].Size.Text)

	for _,v := range metafile.Files.File[0].Verification.Hashes {
		fmt.Printf("%v => %v\n",v.Type, v.Text)
	}
	for _,v := range metafile.Files.File[0].Resources[0].Urls {
		fmt.Printf("%v, %v, %v, %v, %v\n",v.Type, v.Protocol,v.Location,v.Preference,v.Link)
	}
*/

