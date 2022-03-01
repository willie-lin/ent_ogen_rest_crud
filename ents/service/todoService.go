package service

import (
	"context"
	"github.com/willie-lin/ent_ogen_rest_crud/ents/config"
	"github.com/willie-lin/ent_ogen_rest_crud/ents/database/ent"
	"github.com/willie-lin/ent_ogen_rest_crud/ents/database/ent/todo"
	"github.com/willie-lin/ent_ogen_rest_crud/ents/database/ent/user"
)

type TodoOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewTodoOps(ctx context.Context) *TodoOps {
	return &TodoOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *TodoOps) TodoGetByID(id int) (*ent.Todo, error) {
	todo, err := r.client.Todo.Query().WithOwner().Where(todo.ID(id)).Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoOps) TodoCreate(newTodo ent.Todo) (*ent.Todo, error) {
	user, err := r.client.User.Query().Where(user.ID(newTodo.Edges.Owner.ID)).Only(r.ctx)
	if err != nil {
		return nil, err
	}

	todo, err := r.client.Todo.Create().SetName(newTodo.Name).SetOwner(user).Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoOps) UpdateTodo(todo ent.Todo) (*ent.Todo, error) {
	updatedTodo, err := r.client.Todo.UpdateOneID(todo.ID).SetDone(todo.Done).SetName(todo.Name).Save(r.ctx)
	if err != nil {
		return nil, err
	}
	return updatedTodo, nil
}

func (r *TodoOps) DeleteTodo(id int) (int, error) {
	err := r.client.Todo.DeleteOneID(id).Exec(r.ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}
