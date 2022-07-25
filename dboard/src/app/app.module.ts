import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatNativeDateModule } from '@angular/material/core'
import { MaterialModule } from './material.module';

import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { ChartsModule } from '@progress/kendo-angular-charts';

import 'hammerjs';
import { DateInputsModule } from '@progress/kendo-angular-dateinputs';
import { GridModule, ExcelModule   } from '@progress/kendo-angular-grid';
import { ButtonsModule } from "@progress/kendo-angular-buttons";



@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatNativeDateModule,
    MaterialModule,
    HttpClientModule,
    FormsModule, ChartsModule, DateInputsModule, GridModule, ButtonsModule, ExcelModule
  ],
  providers: [],
  bootstrap: [AppComponent],
  exports: [

  ]
})
export class AppModule { }
