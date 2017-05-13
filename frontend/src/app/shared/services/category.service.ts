import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

import { Category, Section } from '../models/category';

import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/publishReplay';

@Injectable()
export class CategoryService {
  baseUrl = 'http://localhost:8181/';
  categories = this.baseUrl + 'categories';
  category = this.baseUrl + 'category';
  sections = this.baseUrl + 'sections';

  constructor(private http: Http) {}

  getSections(): Observable<Section[]> {
    return this.http.get(this.sections)
     .map(response => <Section[]> response.json())
     .catch(this.handleError);
  }

  private handleError (error: Response | any) {
    let errMsg: string;
    if (error instanceof Response) {
      const body = error.json() || '';
      const err = body.error || JSON.stringify(body);
      if (error.status == 0) {
        errMsg = 'failed to contact backend';
      } else {
        errMsg = `${error.status} - ${error.statusText || ''} ${err}`;
      }
    } else {
      errMsg = error.message ? error.message : error.toString();
    }
    return Observable.throw(errMsg);
  }
};