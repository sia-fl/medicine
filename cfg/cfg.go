package cfg

type (
    App struct {
        Debug        bool
        Secret       string
        Database     *Database
        Service      *Service
        AppAdminAddr *Addr
        S3           *S3
        S3Path       *S3Path
        S3Bucket     *S3Bucket
    }

    Addr struct {
        Host string
        Port string
    }

    Service struct {
        Ids *ServiceIds
    }

    S3Path struct {
        CommodityOrigin string
    }

    S3Bucket struct {
        Commodity string
    }

    S3 struct {
        S3ID       string
        S3Secret   string
        Region     string
        Schema     string
        Endpoint   string
        DisableSSL bool
    }
)
