import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { TodoService } from '../todo.service';
import { Todo } from '../todo';

@Component({
  selector: 'app-todos',
  templateUrl: './todos.component.html',
  styleUrls: ['./todos.component.scss']
})
export class TodosComponent implements OnInit {
  
  todos: Todo[] = [];

  message = '';
  umail = '';

  constructor(
    private todoService: TodoService,
    private route: ActivatedRoute) { }

  ngOnInit(): void {
    // this.umail = this.route.snapshot.paramMap.get('email');
    const param = this.route.snapshot.paramMap.get('email');
    // console.log('param',param)
    if(param)this.umail=param
    this.getTodosbyEmail();
    // console.log(umail)
 
    // this.todoService.getTodos()
    //   // .subscribe(td => this.todos = td);
    //   .subscribe({
    //     next: t => this.todos = t,
    //     error: err => this.message = err.message
    //   });
      // console.log("todo",this.todos);
  }

  getTodos(): void {
    this.todoService.getTodos()
    .subscribe(todos => this.todos = todos);
  }

  getTodosbyEmail(): void {
    // console.log(this.umail)
    this.todoService.getTodosbyEmail(this.umail)
    .subscribe(todos => this.todos = todos);
  }

  // onQeury(){
  //   this.todoService.getTodos()
  //     .subscribe({
  //       next: t => this.todos = t,
  //       error: err => this.message = err.message
  //     });
  //     console.log("todo",this.todos);
  // }

  add(todo: string): void {
    // console.log(todo)
    todo = todo.trim();
    // console.log(this.umail)
    // this.onQeury()

    const mytodo: Todo ={
      ID: 0,
      Name: '',
      Email: this.umail,
      Todo: todo
    }
    
    if (!todo) { return; }
    this.todoService.addTodo(mytodo)
      .subscribe(todo => {
        this.todos.push(todo);
        // console.log(this.todos)
        // console.log("sub todo",todo)
        this.getTodosbyEmail()
      });
  }

  delete(todo: Todo): void {
    this.todos = this.todos.filter(t => t !== todo);
    this.todoService.deleteTodo(todo.ID).subscribe();
  }
}


