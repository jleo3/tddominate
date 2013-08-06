---
published: false
layout: post
title: "Retrofitting MVVM: Moving Out"
date: 2013-08-01 08:49
comments: true
categories: 
---

Last week I discovered that I've been lying to myself. I told myself for almost a year that I was making my application conform to the MVVM framework. I was not. Not by a wide margin. I suspected this, then had [StackOverflow confirm it](http://stackoverflow.com/questions/17845668/how-do-i-reduce-dependencies-in-my-bloated-wpf-mainwindow).

After dusting myself off, I decided change was needed. It's one thing to decide against using an MVVM framework, a decision my team made a long time before things got so big and complicated.(My friend and former colleague, Moss Collum, had the best words of wisdom for me when I lamented my architectural state on Twitter: 
@jleo3: 
@moss:

)

 But in our case we've transgressed against known effective patterns of .NET development. Reparations must be made. So I started looking into ways to change. I watched an excellent talk by Rob Eisenberg called ["Build Your Own MVVM Framework"](http://channel9.msdn.com/events/MIX/MIX10/EX15). I read, and I experimented. Then I came up with a plan.

 The first problem comes with our windows. We have three types of windows that the user can load, each with its own nested WinForms grid. One of the views is a "master" view. The other two views can be launched from the master view. One is a detail view and the other is a scenario view. We use AutoFac for dependency injection, but we overloaded our code-behinds with dependencies - mostly ViewModels - instead of using one ViewModel to rule them all.

 I know that a code-behind with lots of dependencies and logic is probably the most egregious of our sins. I also know that moving away from this will give us two advantages: first, it will allow us to have a proper ViewModel per view which can control everything. This sets us up for MVVM retrofitting. Second, extracting this mess by hand will help me to see the common functionality being duplicated across my three windows.

 
