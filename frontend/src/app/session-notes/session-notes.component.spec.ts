import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SessionNotesComponent } from './session-notes.component';

describe('SessionNotesComponent', () => {
  let component: SessionNotesComponent;
  let fixture: ComponentFixture<SessionNotesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [SessionNotesComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(SessionNotesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
