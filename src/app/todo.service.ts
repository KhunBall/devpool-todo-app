import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { MessageService } from './message.service';
import { Todo } from './todo';
// import { TODOS } from './mock-todos'; 

@Injectable({
  providedIn: 'root'
})
export class TodoService {

  private todosUrl = 'http://localhost/users';  // URL to web api
  // private todosUrl = 'http://localhost:5678/users';  // URL to web api

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient,
    private messageService: MessageService) { }

  /** get by mock todos */
  // getTodos(): Observable<Todo[]> {
  //   const todos = of(TODOS)
  //   this.messageService.add('HeroService: fetched heroes');
  //   return todos;
  // }

  /** GET heroes from the server */
  getTodos(): Observable<Todo[]> {
    return this.http.get<Todo[]>(this.todosUrl)
      .pipe(
        tap(_ => this.log('fetched todos')),
        catchError(this.handleError<Todo[]>('gettodos', []))
      );
  }

  getTodosbyEmail(email:string): Observable<Todo[]> {
    const url = `${this.todosUrl}/email/${email}`;
    console.log(url);
    return this.http.get<Todo[]>(url)
      .pipe(
        tap(_ => this.log('fetched todos')),
        catchError(this.handleError<Todo[]>('gettodos', []))
      );
  }

  /** get by mock todos */
  // getTodo(id: number): Observable<Todo> {
  //   // For now, assume that a hero with the specified `id` always exists.
  //   // Error handling will be added in the next step of the tutorial.
  //   const todo = TODOS.find(h => h.ID === id)!;
  //   this.messageService.add(`HeroService: fetched hero id=${id}`);
  //   return of(todo);
  // }

  /** GET hero by id. Will 404 if id not found */
  getTodo(id: number): Observable<Todo> {
    const url = `${this.todosUrl}/${id}`;
    // console.log(url);
    return this.http.get<Todo>(url)
      .pipe(
        // map((data) => {
        //   console.log("mydata",data)
        //    //You can perform some transformation here
        //    return data;
        // }),
        tap(_ => this.log(`fetched hero id=${id}`)),
        catchError(this.handleError<Todo>(`gettodos id=${id}`))
    );
  }

  /** POST: add a new hero to the server */
  addTodo(todo: Todo): Observable<Todo> {
    // console.log("todo", todo)
    return this.http.post<Todo>(this.todosUrl, todo, this.httpOptions).pipe(
      tap((newTodo: Todo) => this.log(`added hero w/ id=${newTodo.ID}`)),
      catchError(this.handleError<Todo>('addHero'))
    );
  }

  /** DELETE: delete the hero from the server */
  deleteTodo(id: number): Observable<Todo> {
    const url = `${this.todosUrl}/${id}`;

    return this.http.delete<Todo>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted todo id=${id}`)),
      catchError(this.handleError<Todo>('deleteHero'))
    );
  }

  /** PUT: update the hero on the server */
  updateTodo(todo: Todo): Observable<any> {
    return this.http.put(this.todosUrl, todo, this.httpOptions).pipe(
      map((data) => {
        console.log("mydata",data)
         //You can perform some transformation here
         return data;
      }),
      tap(_ => this.log(`updated hero id=${todo.ID}`)),
      catchError(this.handleError<any>('updateHero'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   *
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
   private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  /** Log a HeroService message with the MessageService */
  private log(message: string) {
    this.messageService.add(`HeroService: ${message}`);
  }
}
