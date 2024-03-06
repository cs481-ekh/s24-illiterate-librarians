import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { RouterOutlet } from '@angular/router';
import { ProfileParentComponent } from './parent/parent.component';
import { ProfileChildrenComponent } from './children/children.component';
import { ProfileSettingsComponent } from './settings/settings.component';

@Component({
  selector: 'app-profile',
  templateUrl: "./profile.component.html",
  styleUrl: './profile.component.css'
})

export class ProfileComponent { 
  ngOnInit() {
    const items = document.querySelectorAll('.nav .item');
    const parent = document.getElementById('parent');
    const children = document.getElementById('children');
    const settings = document.getElementById('settings');
    const parentComponent = document.querySelector('app-parent');
    const childrenComponent = document.querySelector('app-children');
    const settingsComponent = document.querySelector('app-settings');

    if (parent && children && settings && parentComponent && childrenComponent && settingsComponent) {
      const parentComponent = document.querySelector('app-parent') as HTMLElement;
      const childrenComponent = document.querySelector('app-children') as HTMLElement;
      const settingsComponent = document.querySelector('app-settings') as HTMLElement;

      parent.addEventListener('click', function () {
        parentComponent.style.display = 'block';
        childrenComponent.style.display = 'none';
        settingsComponent.style.display = 'none';
      });

      children.addEventListener('click', function () {
        parentComponent.style.display = 'none';
        childrenComponent.style.display = 'block';
        settingsComponent.style.display = 'none';
      });

      settings.addEventListener('click', function () {
        parentComponent.style.display = 'none';
        childrenComponent.style.display = 'none';
        settingsComponent.style.display = 'block';
      });
    }
  }
}