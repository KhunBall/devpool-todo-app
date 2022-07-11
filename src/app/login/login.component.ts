import { Component, OnInit } from '@angular/core';
import { Todo } from '../todo';
import { TodoService } from '../todo.service';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  todos: Todo[] = [];

  constructor(
    private todoService: TodoService,
    private route: ActivatedRoute,
    private router: Router) { }

  ngOnInit(): void {
    this.getTodos();
  }

  getTodos(): void {
    this.todoService.getTodos()
      .subscribe(todos => this.todos = todos.slice(1, 5));
  }
  // { path: 'todos/email/:email', component: TodosComponent },
  login(uemail: string): void {
    // console.log(uemail)
    // this.router.navigate(['todos/', {email:uemail}]);
    this.router.navigate(['todos',uemail])
  //   todo = todo.trim();
  //   this.onQeury()
  //   if (!todo) { return; }
  //   this.todoService.addTodo({ todo } as unknown as Todo)
  //     .subscribe(todo => {
  //       this.todos.push(todo);
  //     });
  }

}
