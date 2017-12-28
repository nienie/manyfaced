package source

import (
    "database/sql"
    "fmt"
    "time"

    //neglect it~
    _ "github.com/go-sql-driver/mysql"
    "github.com/didi/gendry/scanner"
    "github.com/nienie/manyfaced/poll"
)

//DBConfigurationSource ...
type DBConfigurationSource struct {
    DBSource        *sql.DB
    QuerySQL        string
    KeyColumnName   string
    ValueColumnName string
}

//NewDBConfigurationSource ...
func NewDBConfigurationSource(dbSource *sql.DB, querySQL string, keyColumnName string, valueColumnName string)*DBConfigurationSource {
    return &DBConfigurationSource{
        DBSource:           dbSource,
        QuerySQL:           querySQL,
        KeyColumnName:      keyColumnName,
        ValueColumnName:    valueColumnName,
    }
}

//Poll ...
func (s *DBConfigurationSource)Poll(initial bool, checkPoint interface{}) (*poll.PolledResult, error) {
    result, err := s.load()
    return poll.NewFullPolledResult(result), err
}

func (s *DBConfigurationSource)load() (map[string]interface{}, error) {
    if err := s.DBSource.Ping(); err != nil {
        return nil, err
    }
    rows, err := s.DBSource.Query(s.QuerySQL)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    records, err := scanner.ScanMap(rows)
    if err != nil {
        return nil, err
    }
    result := make(map[string]interface{})
    for _, record := range records {
        key := fmt.Sprint(formatDBInterface(record[s.KeyColumnName]))
        val := formatDBInterface(record[s.ValueColumnName])
        result[key] = val
    }
    return result, nil
}

func formatDBInterface(v interface{}) interface{} {
    //See: https://golang.org/pkg/database/sql/driver/#Value
    switch v.(type) {
    case int64:
        return v
    case float64:
        return v
    case bool:
        return v
    case []byte:
        return string(v.([]byte))
    case time.Time:
        return fmt.Sprint(v)
    default:
        return v
    }
}





