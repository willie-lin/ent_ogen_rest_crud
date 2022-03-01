// Code generated by entc, DO NOT EDIT.

package ogent

import "github.com/willie-lin/ent_ogen_rest_crud/ents/database/ent"

func NewTodoCreate(e *ent.Todo) *TodoCreate {
	if e == nil {
		return nil
	}
	return &TodoCreate{
		ID:   e.ID,
		Name: e.Name,
		Done: NewOptBool(e.Done),
	}
}

func NewTodoCreates(es []*ent.Todo) []TodoCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]TodoCreate, len(es))
	for i, e := range es {
		r[i] = NewTodoCreate(e).Elem()
	}
	return r
}

func (t *TodoCreate) Elem() TodoCreate {
	if t != nil {
		return TodoCreate{}
	}
	return *t
}

func NewTodoList(e *ent.Todo) *TodoList {
	if e == nil {
		return nil
	}
	return &TodoList{
		ID:   e.ID,
		Name: e.Name,
		Done: NewOptBool(e.Done),
	}
}

func NewTodoLists(es []*ent.Todo) []TodoList {
	if len(es) == 0 {
		return nil
	}
	r := make([]TodoList, len(es))
	for i, e := range es {
		r[i] = NewTodoList(e).Elem()
	}
	return r
}

func (t *TodoList) Elem() TodoList {
	if t != nil {
		return TodoList{}
	}
	return *t
}

func NewTodoRead(e *ent.Todo) *TodoRead {
	if e == nil {
		return nil
	}
	return &TodoRead{
		ID:   e.ID,
		Name: e.Name,
		Done: NewOptBool(e.Done),
	}
}

func NewTodoReads(es []*ent.Todo) []TodoRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]TodoRead, len(es))
	for i, e := range es {
		r[i] = NewTodoRead(e).Elem()
	}
	return r
}

func (t *TodoRead) Elem() TodoRead {
	if t != nil {
		return TodoRead{}
	}
	return *t
}

func NewTodoUpdate(e *ent.Todo) *TodoUpdate {
	if e == nil {
		return nil
	}
	return &TodoUpdate{
		ID:   e.ID,
		Name: e.Name,
		Done: NewOptBool(e.Done),
	}
}

func NewTodoUpdates(es []*ent.Todo) []TodoUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]TodoUpdate, len(es))
	for i, e := range es {
		r[i] = NewTodoUpdate(e).Elem()
	}
	return r
}

func (t *TodoUpdate) Elem() TodoUpdate {
	if t != nil {
		return TodoUpdate{}
	}
	return *t
}

func NewTodoOwnerRead(e *ent.User) *TodoOwnerRead {
	if e == nil {
		return nil
	}
	return &TodoOwnerRead{
		ID:    e.ID,
		Email: e.Email,
		Name:  e.Name,
	}
}

func NewTodoOwnerReads(es []*ent.User) []TodoOwnerRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]TodoOwnerRead, len(es))
	for i, e := range es {
		r[i] = NewTodoOwnerRead(e).Elem()
	}
	return r
}

func (u *TodoOwnerRead) Elem() TodoOwnerRead {
	if u != nil {
		return TodoOwnerRead{}
	}
	return *u
}

func NewUserCreate(e *ent.User) *UserCreate {
	if e == nil {
		return nil
	}
	return &UserCreate{
		ID:    e.ID,
		Email: e.Email,
		Name:  e.Name,
	}
}

func NewUserCreates(es []*ent.User) []UserCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserCreate, len(es))
	for i, e := range es {
		r[i] = NewUserCreate(e).Elem()
	}
	return r
}

func (u *UserCreate) Elem() UserCreate {
	if u != nil {
		return UserCreate{}
	}
	return *u
}

func NewUserList(e *ent.User) *UserList {
	if e == nil {
		return nil
	}
	return &UserList{
		ID:    e.ID,
		Email: e.Email,
		Name:  e.Name,
	}
}

func NewUserLists(es []*ent.User) []UserList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserList, len(es))
	for i, e := range es {
		r[i] = NewUserList(e).Elem()
	}
	return r
}

func (u *UserList) Elem() UserList {
	if u != nil {
		return UserList{}
	}
	return *u
}

func NewUserRead(e *ent.User) *UserRead {
	if e == nil {
		return nil
	}
	return &UserRead{
		ID:    e.ID,
		Email: e.Email,
		Name:  e.Name,
	}
}

func NewUserReads(es []*ent.User) []UserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserRead, len(es))
	for i, e := range es {
		r[i] = NewUserRead(e).Elem()
	}
	return r
}

func (u *UserRead) Elem() UserRead {
	if u != nil {
		return UserRead{}
	}
	return *u
}

func NewUserUpdate(e *ent.User) *UserUpdate {
	if e == nil {
		return nil
	}
	return &UserUpdate{
		ID:    e.ID,
		Email: e.Email,
		Name:  e.Name,
	}
}

func NewUserUpdates(es []*ent.User) []UserUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserUpdate, len(es))
	for i, e := range es {
		r[i] = NewUserUpdate(e).Elem()
	}
	return r
}

func (u *UserUpdate) Elem() UserUpdate {
	if u != nil {
		return UserUpdate{}
	}
	return *u
}

func NewUserTodosList(e *ent.Todo) *UserTodosList {
	if e == nil {
		return nil
	}
	return &UserTodosList{
		ID:   e.ID,
		Name: e.Name,
		Done: NewOptBool(e.Done),
	}
}

func NewUserTodosLists(es []*ent.Todo) []UserTodosList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserTodosList, len(es))
	for i, e := range es {
		r[i] = NewUserTodosList(e).Elem()
	}
	return r
}

func (t *UserTodosList) Elem() UserTodosList {
	if t != nil {
		return UserTodosList{}
	}
	return *t
}
