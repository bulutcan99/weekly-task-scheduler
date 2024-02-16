# Weekly Dev Task Scheduler

Welcome to the Weekly Dev Task Scheduler! This Go project implements a task distributor utilizing the Domain-Driven Design (DDD) approach. It efficiently distributes tasks fetched from various providers to developers on a weekly basis. Our aim is to provide a clear and organized overview of tasks, prioritizing them based on estimated duration and value, ensuring that developers are assigned tasks that are best suited to their schedule and expertise.

## Features

-   **Task Synchronization**: Sync tasks with developers for fastest way possible.
-   **Provider Filtering**: View tasks for a specific provider, or see an aggregated list.
-   **Domain-Driven Design (DDD):** Embraces DDD principles for a clean and maintainable codebase.
-   **Provider Integration:** Fetches tasks from multiple providers seamlessly.
-   **Weekly Task Distribution:** Automatically distributes tasks to developers on a weekly schedule.
-   **Scalable:** Designed to handle a growing number of tasks and developers.

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

-   Docker
-   Make (Windows: https://gnuwin32.sourceforge.net/packages/make.htm, Mac: brew install make)


## Setup

-   Clone the repository
-   Copy `.env.example` to `.env` and fill in your credentials
-   Run `make docker-install`
-   Run `make run-seeder` // for seed the two mock json datas
-   Run `make run-server`


## Usage

-   Go to `http://localhost:8080/swagger/index.html#/` main page for swagger
-   Click on post provider endpoint to add providers
-   Example provider json responses are :

    -   `https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd` for provider 1
    -   `https://run.mocky.io/v3/7b0ff222-7a9c-4c54-9396-0df58e289143` for provider 2

-   They can add any kind of data as long as their types are suitable.
