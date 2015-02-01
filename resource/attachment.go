package resource

import(
	"io/ioutil"
	"encoding/base64"
	"bytes"
	. "github.com/talbright/go-desk/types"
)

type Attachment struct {
	Size            *int                   `json:"size,omitempty"`
	FileName        *string                `json:"file_name,omitempty"`
	ContentType     *string                `json:"content_type,omitempty"`
	Content         *string                `json:"content,omitempty"`
	URL             *string                `json:"url,omitempty"`
	ErasedAt        *Timestamp             `json:"erased_at,omitempty"`
	CreatedAt       *Timestamp             `json:"created_at,omitempty"`
	UpdatedAt       *Timestamp             `json:"updated_at,omitempty"`
	Resource
}

func NewAttachment() *Attachment {
	attachment := &Attachment{}
	attachment.InitializeResource(attachment)
	return attachment
}

func (r *Attachment) SetContent(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err == nil {
		buf := new(bytes.Buffer)
		encoder := base64.NewEncoder(base64.StdEncoding,buf)
		encoder.Write(data)
		encoder.Close()
		bytes := buf.Bytes()
		r.Content = String(string(bytes[:]))
	}
	return err
}

func (r Attachment) String() string {
	return Stringify(r)
}
