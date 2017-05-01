import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

import { Category } from '../models/category';

import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/publishReplay';

@Injectable()
export class CategoryService {
  baseUrl = 'http://localhost:8181/';
  categories = this.baseUrl + 'categories';
  category = this.baseUrl + 'category';

  constructor(private http: Http) {}

  getCategories(): Observable<Category[]> {
    return this.http.get(this.categories)
     .map(response => <Category[]> response.json())
     .catch(this.handleError);
  }

  getCategory(id): Observable<Category> {
    return this.http.get(this.category + '/' + id)
      .map(response => <Category> response.json())
      .catch(this.handleError)
      .publishReplay(1)
      .refCount()
  }

  private handleError (error: Response | any) {
    let errMsg: string;
    if (error instanceof Response) {
      const body = error.json() || '';
      const err = body.error || JSON.stringify(body);
      errMsg = `${error.status} - ${error.statusText || ''} ${err}`;
    } else {
      errMsg = error.message ? error.message : error.toString();
    }
    console.error(errMsg);
    return Observable.throw(errMsg);
  }
};