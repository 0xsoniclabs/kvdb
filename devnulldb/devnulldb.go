package devnulldb

import (
	"github.com/0xsoniclabs/kvdb"
)

// Database is an always empty database.
type Database struct{}

// New returns an always empty database.
func New() *Database {
	return &Database{}
}

// Close deallocates the internal map and ensures any consecutive data access op
// failes with an error.
func (db *Database) Close() error {
	return nil
}

// Drop whole database.
func (db *Database) Drop() {
}

// Has retrieves if a key is present in the key-value store.
func (db *Database) Has(key []byte) (bool, error) {
	return false, nil
}

// Get retrieves the given key if it's present in the key-value store.
func (db *Database) Get(key []byte) ([]byte, error) {
	return nil, nil
}

// Put inserts the given value into the key-value store.
func (db *Database) Put(key []byte, value []byte) error {
	return nil
}

// Delete removes the key from the key-value store.
func (db *Database) Delete(key []byte) error {
	return nil
}

// NewBatch creates a write-only key-value store that buffers changes to its host
// database until a final write is called.
func (db *Database) NewBatch() kvdb.Batch {
	return &batch{}
}

// NewIterator creates a binary-alphabetical iterator over a subset
// of database content with a particular key prefix, starting at a particular
// initial key (or after, if it does not exist).
func (db *Database) NewIterator(prefix []byte, start []byte) kvdb.Iterator {
	return &iterator{}
}

// GetSnapshot returns a latest snapshot of the underlying DB. A snapshot
// is a frozen snapshot of a DB state at a particular point in time. The
// content of snapshot are guaranteed to be consistent.
//
// The snapshot must be released after use, by calling Release method.
func (db *Database) GetSnapshot() (kvdb.Snapshot, error) {
	return &Snapshot{New()}, nil
}

// Stat returns a particular internal stat of the database.
func (db *Database) Stat() (string, error) {
	return "", nil
}

// Compact is not supported on a memory database.
func (db *Database) Compact(start []byte, limit []byte) error {
	return nil
}

// Len returns the number of entries currently present in the memory database.
//
// Note, this method is only used for testing (i.e. not public in general) and
// does not have explicit checks for closed-ness to allow simpler testing code.
func (db *Database) Len() int {
	return 0
}

func (db *Database) AncientDatadir() (string, error) {
	return "", nil
}

// batch is a write-only memory batch that commits changes to its host
// database when Write is called. A batch cannot be used concurrently.
type batch struct{}

// Put inserts the given value into the batch for later committing.
func (b *batch) Put(key, value []byte) error {
	return nil
}

// Delete inserts the a key removal into the batch for later committing.
func (b *batch) Delete(key []byte) error {
	return nil
}

// ValueSize retrieves the amount of data queued up for writing.
func (b *batch) ValueSize() int {
	return 0
}

// Write flushes any accumulated data to the memory database.
func (b *batch) Write() error {
	return nil
}

// Reset resets the batch for reuse.
func (b *batch) Reset() {
}

// Replay replays the batch contents.
func (b *batch) Replay(w kvdb.Writer) error {
	return nil
}

// DeleteRange deletes a range of the batch.
func (b *batch) DeleteRange(start, end []byte) error {
	return nil
}

// iterator can walk over the (potentially partial) keyspace of a memory key
// value store. Internally it is a deep copy of the entire iterated state,
// sorted by keys.
type iterator struct{}

// Next moves the iterator to the next key/value pair. It returns whether the
// iterator is exhausted.
func (it *iterator) Next() bool {
	return false
}

// Error returns any accumulated error. Exhausting all the key/value pairs
// is not considered to be an error. A memory iterator cannot encounter errors.
func (it *iterator) Error() error {
	return nil
}

// Key returns the key of the current key/value pair, or nil if done. The caller
// should not modify the contents of the returned slice, and its contents may
// change on the next call to Next.
func (it *iterator) Key() []byte {
	return nil
}

// Value returns the value of the current key/value pair, or nil if done. The
// caller should not modify the contents of the returned slice, and its contents
// may change on the next call to Next.
func (it *iterator) Value() []byte {
	return nil
}

// Release releases associated resources. Release should always succeed and can
// be called multiple times without causing error.
func (it *iterator) Release() {
}

// Snapshot is a DB snapshot.
type Snapshot struct {
	*Database
}

func (s *Snapshot) Release() {
	s.Database = nil
}
