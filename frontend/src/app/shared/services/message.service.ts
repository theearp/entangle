import { Injectable } from '@angular/core';

import { Observable } from 'rxjs/Observable';
import { Subject } from 'rxjs';

@Injectable()
export class MessageService {
  private notification: Subject<string> = new Subject();
  constructor() { }

  send(msg: string) {
    this.notification.next(msg);
  }

  clear() {
    this.notification.next();
  }

  get(): Observable<string> {
    return this.notification.asObservable();
  }
}