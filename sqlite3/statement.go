package sqlite3

import "C"

func (s *Stmt) BindString(v string, index int) error {
	// the leftmost SQL parameter has an index of 1.
	rc := C.bind_text(s.stmt, C.int(index+1), cStr(v), C.int(len(v)), 1)
	if rc != OK {
		return errStr(rc)
	}
	return nil
}
