package pdnsjsonrpc

import (
	"net"
)

type InitializeQuery map[string]interface{}

type BoolResponse bool

type LookupQuery struct {
	QType      string  `json:"qtype"`
	QName      string  `json:"qname"`
	ZoneId     int     `json:"zone-id"`
	Remote     *net.IP `json:"remote,omitempty"`
	Local      *net.IP `json:"local,omitempty"`
	RealRemote *net.IP `json:"real-remote,omitempty"`
}

type LookupResponseEntry struct {
	QType     string `json:"qtype"`
	QName     string `json:"qname"`
	Content   string `json:"content"`
	Ttl       int    `json:"ttl"`
	DomainId  int    `json:"domain_id,omitempty"`
	ScopeMask int    `json:"scopeMask,omitempty"`
	Auth      int    `json:"auth,omitempty"`
}

type LookupResponse []LookupResponseEntry

type ListQuery struct {
	ZoneName string `json:"zonename"`
	DomainId int    `json:"domain_id,omitempty"`
}

// List returns a LookupResponse

type BeforeAndAfterQuery struct {
	Id    int    `json:"id"`
	QName string `json:"qname"`
}

type BeforeAndAfterResponse struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type AllDomainMetadataQuery struct {
	Name string `json:"name"`
}

type AllDomainMetadataResponse map[string][]string

type DomainMetadataQuery struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type DomainMetadataResponse []string

type DomainKeysQuery struct {
	Name string `json:"name"`
	Kind string `json:"kind,omitempty"`
}

type DomainKeysResponseEntry struct {
	Id      int    `json:"id"`
	Flags   int    `json:"flags"`
	Active  bool   `json:"active"`
	Content string `json:"content"`
}

type DomainKeysResponse []DomainKeysResponseEntry

type AddDomainKeyQueryKey struct {
	Flags   int    `json:"flags"`
	Active  bool   `json:"active"`
	Content string `json:"content"`
}

type AddDomainKeyQuery struct {
	Name string               `json:"name"`
	Key  AddDomainKeyQueryKey `json:"key"`
	Id   int                  `json:"id"`
}

// Returns BoolResponse

type DomainKeyNameIdPair struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type RemoveDomainKeyQuery DomainKeyNameIdPair

// Returns BoolResponse

type ActivateDomainKeyQuery DomainKeyNameIdPair

// Returns BoolResponse

type DeactivateDomainKeyQuery DomainKeyNameIdPair

type GetTSigKeyQuery struct {
	Name string `json:"name"`
}

type GetTSigKeyResponse struct {
	Algorithm string `json:"algorithm"`
	Content   string `json:"content"`
}

type SetNotifiedQuery struct {
	Id     int `json:"id"`
	Serial int `json:"serial"`
}

// Returns BoolResponse

type IsMasterQuery struct {
	Name string `json:"name"`
	Ip   net.IP `json:"ip"`
}

// Returns BoolResponse

type SuperMasterBackendQuery struct {
	Ip      net.IP                `json:"ip"`
	Domain  string                `json:"domain"`
	Nsset   []LookupResponseEntry `json:"nsset"`
	Account string                `json:"account"`
}

// Can also return a BoolResponse
type SuperMasterBackendResponse struct {
	Account    string `json:"account"`
	Nameserver string `json:"nameserver"`
}

type CreateSlaveDomainQuery struct {
	Ip         net.IP `json:"ip"`
	Domain     string `json:"domain"`
	Nameserver string `json:"nameserver,omitempty"`
	Account    string `json:"account,emitempty"`
}

// Returns BoolResponse

type ReplaceRRSetQuery struct {
	DomainId int                   `json:"domain_id"`
	QName    string                `json:"qname"`
	QType    string                `json:"qtype"`
	TrxId    int                   `json:"trxid"`
	RRSet    []LookupResponseEntry `json:"rrset"`
}

// Returns BoolResponse

type FeedRecordQuery struct {
	RR    LookupResponseEntry `json:"rr"`
	TrxId int                 `json:"trxid"`
}

// Returns BoolResponse

type FeedEntsQuery struct {
	DomainId int      `json:"domain_id"`
	TrxId    int      `json:"trxid"`
	Nonterm  []string `json:"nonterm"`
}

// Returns BoolResponse

type FeedEnts3Query struct {
	DomainId int      `json:"domain_id"`
	Domain   string   `json:"domain"`
	TrxId    int      `json:"trxid"`
	Times    int      `json:"times"`
	Salt     string   `json:"salt"`
	Narrow   bool     `json:"narrow"`
	Nonterm  []string `json:"nonterm"`
}

// Returns BoolResponse

type StartTransactionQuery struct {
	DomainId int    `json:"domain_id"`
	Domain   string `json:"domain"`
	TrxId    int    `json:"trxid"`
}

// Returns BoolResponse

type TransactionQuery struct {
	TrxId int `json:"trxid"`
}

type CommitTransactionQuery TransactionQuery
type AbortTransactionQuery TransactionQuery

type CalculateSOASerialQuery struct {
	Domain string                 `json:"domain"`
	Sd     map[string]interface{} `json:"sd"` // I really don't know what is in here
}

type CalculateSOASerialResponse int

type DirectBackendCmdQuery struct {
	Query string `json:"query"`
}

type DirectBackendCmdResponse string

type GetAllDomainsQuery struct {
	IncludeDisabled bool `json:"include_disabled"`
}

type GetAllDomainsResponse []map[string]interface{}

type SearchRecordsQuery struct {
	Pattern    string `json:"pattern"`
	MaxResults int    `json:"maxResults"`
}

// Can also be false
type SearchResultsResponse []LookupResponseEntry
