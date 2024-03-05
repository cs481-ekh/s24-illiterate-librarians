import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfileChildrenComponent } from './children.component';

describe('ChildrenComponent', () => {
  let component: ProfileChildrenComponent;
  let fixture: ComponentFixture<ProfileChildrenComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProfileChildrenComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ProfileChildrenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
