---
title: Let's build
date: 31.03.2024
---

So, the tech.
Oh yes, the most important part — the tech.
I'm going to use stuff I'm most comfortable with, which is happened to be the most widespread tech stack in the world:
Angular frontend, Java + String backend, and all that on top of AWS.

Let's begin with infrastructure — to keep things simple, I'm using AWS Amplify to run frontend, and AWS AppRunner to run
backend.
For now, there's no need for anything more complex than this.

### AWS Amplify

AWS Amplify hooks up to frontend repository via GitHub webhook.
And everytime anything is pushed into `main` branch, Amplify gets notified and the CI/CD machinery kicks in.
Amplify is smart enough to understand that it's connected to angular app (this actually doesn't matter,
because it builds a project with a silly `npm run build` script).

Build artifact is then stored into AWS S3 bucket
(unfortunately, or not, this bucket is not accessible)
and then exposed via cloudfront distribution(also not accessible).
By "not accessible" I mean that it's not created under my account, I can't look at it nor touch it.
It exists, but somewhere within the bowels of AWS.
Serverless, right?

AWS S3 is a perfect place for frontend artifacts – infinitely scalable, ultimately robust, publicly accessible(when
needed), cheap.
It just works.
I have a strong impression that AWS S3 powers at least half of the internet,
and so I'm trusting it to host my amazing frontend.

### AWS AppRunner

After the first blog post, I had no backend for my blog application.

— "Do I even need a backend?" - was my question.

— Of course, I'm a backend developer, I have to have a backend.

— Alright, let's have it.

Building backend is straightforward.
Code here, code there — I'm doing this for the last 15 years, so I'm feeling somewhat comfortable.
The real question is "How to run it?"

EKS?
Hell no, I'm not touching Kubernetes.
I'm sick of it.
It's too complex.
Moreover, I want to run a single container.
To say that EKS is an overkill in this situation is a huge understatement.

ECS?
Sounds better.
Let's do it.
I've created a cluster, task definition, created a task... and nothing.
I can't access my service from the outside.
Oh, no... networking.
Something is not right with the VPC setup.
Subset seems fine.
Security groups and routing tables are also "looks fine."
Damn it, something silly is not right, and I can't find it.
Screw it — a task stopped, task definition deleted, cluster deleted.
ECS is also too complex.

While in bed and half asleep, I was browsing through AWS Console app on my phone.
Eureka!
AWS Q. AWS AI assistant.
This is exactly what they built it for — so that idiots like me could ask questions like mine.
The answer was instant — AWS AppRunner.
Next morning I logged in into AWS AppRunner, clicked a few buttons,
selected a hello world image from ECR, "deploy" and... it worked.
My hello world backend is running in a matter of 2–3 minutes.
This is why I love AWS.

I've hidden my app deployment via a custom domain `api.buyallmemes.com` by fiddling with Route 53 hosted zone.
Thankfully, I know a couple of tricks around DNS.

Now, it's time to build the real backend.

### Java + Spring = ❤️

The choice of tech for the backend is super easy. There's no choice really.
There's only one true kind, and it's Java + Spring.
I'm starting with an extremely simple setup: one REST endpoint that returns a list of posts.
What is a post?
A simple resource with only one attribute — content.
For now, I don't need anything else.

However, I do need something — Zalando Problem library https://github.com/zalando/problem.
I'm sure you're aware of Zalando as an internet cloth retailer, but you might not be aware that they have quite a few
cool bits of software.
Problem Library is one of those bits.
It's a small library with a single purpose — unify an approach for expressing errors in REST API.
Instead of figuring out every time what to return in case of error,
or returning gibberish (like a full Spring Web stack stace in case of 500),
zalando/problem library suggests returning their little `Problem` structure.
Naturally, a library has an awesome integration with Spring, so there's very little configuration required.
Use it, do yourself (and your REST API consumers) a favor.

Another one of those hidden gems is a Zalando RESTful API
Guidelines https://opensource.zalando.com/restful-api-guidelines/ — read it.
It's awesome.

So, after the initial setup, I throw a bunch of code in.

**Rule #1: First, make it work, then make it right, then make it fast.**

I don't care about performance at the moment(if ever), so I will ignore the latter part.
Let's focus on making things work.

Damn it, I need a database to store posts!
Or do I?
Hmm, why the hell would I need an enterprise grade DB (like PostgreSQL) to store a single post - sounds absurd.
I will store it on disk as part of the source code!
My IDE is a perfect .MD editor.
Git will provide me with all the version control I ever need.
I can just branch out of the `main`, write whatever I want, and then merge it back when it's ready to be published.
And it's free!

Well, I need to redeploy the backend every time I write or change the post,
but for now, this is not a big deal, so this mechanism will suffice.
I've set AWS AppRunner to automatically detect and deploy the newest image versions of my backend.
So I don't have to do much manual stuff, besides building an image.

Btw, how do I suppose to build and push image into ECR?
I'm not writing Dockerfile — that's for sure.
Google Jib, https://github.com/GoogleContainerTools/jib.

Simple jib gradle plugin declaration in `build.gradle`(Gradle FTW!),
set `jib.from.image` parameter to `amazoncorretto:21-alpine`, set `jib.to.image` to my ECR repo.
Quick `aws ecr get-login-password...` from ECR documentation, `./gradlew jib` and off flies my images.
Easy enough.
I will automate it later.
I think GitHub Actions is what cool kids are using (I'm more of a GitLab user,
but for the sake of exercise, I decided to publish everything on GitHub).

Alright, for now, that's enough.
I have a running Angular frontend and Java backend.
Frontend knows how to talk with backend.
Backend return list of posts, which are store in `resources` folder.
Backend logic is rather silly

- Read files from `resources/blog/posts` project folder
- Load each file content as string into a Post object
- Sort loaded posts by filename in descending order

And yes, I've introduced `fileName` attribute to the `Post`.
And that's about it.
I already established minimal flow of work.

At the moment, there's little to talk about.
There's little code and one cute unit test.
I guess this is worth talking about — I'm a huge fan of TDD.
I love my tests.
At the moment, I have only one but crucial test that covers two most important aspects — REST endpoint and that posts
are properly ordered.
I decided to use file naming as a sort parameter.
Each new post-file will be prefixed by the current date,
so I could easily sort them in reverse order to show the latest posts on top, and oldest at the bottom.
Since I'm a backend guy, I prefer to keep such logic at the back.
I don't want to spend much time on frontend, so I will try to keep it as lean as possible.
Saying that, the more I think about it, the more I realize that I should've gone with something like a thymeleaf,
and build everything within the backend app, but what's done is done.
Having a separate frontend app is not without its benefits anyway.
Plus, I can definitely benefit from expanding my horizons beyond backend and Java.



