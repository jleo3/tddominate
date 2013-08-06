---
layout: post
title: "Use Automocking to Mop Up Your Specs"
date: 2013-08-02 15:17
comments: true
categories: 
---

I'm big on mocking. I create loads of interfaces and I test up against my interfaces. That's how I define unit tests. I understand there are plenty of other definitions out there, and many of them have merit. This one helps me to stay organized.

This also leads to loads of setup in my MSpec files:

{% gist 6169321 %}

This class is my setup. My other specs extend this class, reducing some of the noise:

```
    [Subject("CurrentScenario")]
    internal class when_setting_current_scenario : ScenarioWindowViewModelSpecs
    {
        const string CurrentScenario = "current scenario";

        Because of = () => ScenarioWindowViewModel.CurrentScenario = CurrentScenario;

        It will_set_the_current_scenario = () => 
            ScenarioWindowViewModel.CurrentScenario.ShouldContain(CurrentScenario);
    }
```

So, nightmare setup but a pretty concise spec. The problem is compounded, of course, anytime I add, remove, or change my dependencies. I always have to change them in both production and test code.

Now enter automocking, an intelligent mix of unit testing and dependency injection. Using MSpec's automocking capabilities, we declare our Subject Under Test (SUT) and get right to business:

```
    [Subject("CurrentScenario")]
    internal class when_getting_setting : WithSubject<ScenarioWindowViewModel>
    {
        private const string CurrentScenario = "current scenario";

        private Because of = () => Subject.CurrentScenario = CurrentScenario;

        private It will_set_the_current_scenario = () => 
            Subject.CurrentScenario.ShouldContain(CurrentScenario);
    }
```

Where did the setup go? Using this method, it's all done for us. MSpec knows from the class declaration that ScenarioWindowViewModel is our SUT and that there are four interface dependencies. So we don't need to go through the trouble of declaring and intitializing each one.

We don't even have to decare our SUT. We use the `subject` keyword in place of ScenarioWindowViewModel.

This helps tremendously with reducing setup and headaches as a result of churn in the production code. It even provides a nicer syntax for when we need those dependencies. For example, our SUT has a dependency on an IGridSettingsViewModel. Here's how to set an expectation:

```
		It will_register_a_column_manager = () => The<IGridSettingsViewModel>.WasToldTo(
                x => x.SetColumnStyleManager(Arg<IColumnStyleManager>.Is.Anything));
```

Nice and easy, and much cleaner. 
