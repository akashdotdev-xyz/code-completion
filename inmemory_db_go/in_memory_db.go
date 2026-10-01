package db

import (
	"fmt"
	"slices"
)

// InMemoryDB: see README.md for the full spec.
// A nil pointer return means "not found".

type Field struct {
	ID        string
	key       string
	value     string
	expiresAt int
	isExpired bool
}

func (f *Field) GetField(t int, key string) *string {
	if f.isExpired {
		return nil
	}

	if f.expiresAt >= t {
		f.isExpired = true
		return nil
	}

	return &f.value
}

type Record struct {
	ID     string
	fields map[string]*Field
}

type InMemoryDB struct {
	data        map[string]*Record
	expiryQueue []*Field
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]*Record),
	}
}

// ---------- Level 1 ----------

func (d *InMemoryDB) Set(timestamp int, key, field, value string) {
	record, ok := d.data[key]
	if !ok {
		record = &Record{
			ID:     key,
			fields: map[string]*Field{},
		}
	}
	record.fields[field] = &Field{key: field, value: value}
	d.data[key] = record
}

func (d *InMemoryDB) Get(timestamp int, key, field string) *string {
	record, ok := d.data[key]
	if !ok {
		return nil
	}

	fieldData, ok := record.fields[field]
	if !ok {
		return nil
	}
	return fieldData.GetField(timestamp, field)

}

func (d *InMemoryDB) Delete(timestamp int, key, field string) bool {
	val := d.Get(timestamp, key, field)
	if val == nil {
		return false
	}

	delete(d.data[key].fields, field)
	return true
}

// ---------- Level 2 ----------

func (d *InMemoryDB) Scan(timestamp int, key string) []string {
	result := []string{}
	record, ok := d.data[key]
	if !ok {
		return result
	}

	for key, value := range record.fields {
		result = append(result, key+fmt.Sprintf("(%s)", value.value))
	}

	slices.Sort(result)
	return result
}

func (d *InMemoryDB) ScanByPrefix(timestamp int, key, prefix string) []string {

	result := d.Scan(timestamp, key)
	finalResult := []string{}
	for _, str := range result {
		if str[0:len(prefix)] == prefix {
			finalResult = append(finalResult, str)
		}
	}
	return finalResult
}

// ---------- Level 3 ----------

func (d *InMemoryDB) cleanUpExpiredFields(timestamp int) {
	i := 0
	for i < len(d.expiryQueue) && d.expiryQueue[i].expiresAt >= timestamp {
		fieldData := d.expiryQueue[i]

		d.expiryQueue = d.expiryQueue[1:]
		fieldData.isExpired = true
	}
}

func (d *InMemoryDB) SetWithTTL(timestamp int, key, field, value string, ttl int) {
	d.cleanUpExpiredFields(timestamp)
	record, ok := d.data[key]
	if !ok {
		record = &Record{
			ID:     key,
			fields: map[string]*Field{},
		}
	}
	fieldData := &Field{ID: key, key: field, value: value, expiresAt: timestamp + ttl}
	d.expiryQueue = append(d.expiryQueue, fieldData)

	record.fields[field] = fieldData
	d.data[key] = record
}

// ---------- Level 4 ----------

func (d *InMemoryDB) Backup(timestamp int) int {
	return 0
}

func (d *InMemoryDB) Restore(timestamp int, timestampToRestore int) {
}
