package supplier

import s3Session "github.com/aws/aws-sdk-go/aws/session"

type Supplier struct {
    s3sess *s3Session.Session
}

func New(s3sess *s3Session.Session) *Supplier {
    return &Supplier{s3sess: s3sess}
}
