---
name: canvas-cli
description: >
  Use when the user wants to interact with Canvas LMS via the CLI — managing courses, assignments,
  enrollments, users, grades, submissions, or any other Canvas resource. Trigger whenever the user
  mentions "canvas-cli", asks to list/get/create/update/delete any Canvas resource, wants to run a
  Canvas API operation, needs to set up auth, or asks what CLI commands are available. Also trigger
  for the compound commands: student-pulse, course-audit, grade-export, bulk-enroll, activity-report.
author: "simple-scalable-solutions"
license: "Apache-2.0"
argument-hint: "<command> [args]"
allowed-tools: "Read Bash"
---

# Canvas CLI

## Prerequisites: Install the CLI

This skill drives the `canvas-cli` binary. Verify it is installed before running any command:

```bash
canvas-cli --version
```

If missing, install it:

```bash
curl -sSfL https://raw.githubusercontent.com/simple-scalable-solutions/canvas-cli/main/install.sh | sh
```

Then set up credentials (one time only):

```bash
canvas-cli auth set-url https://canvas.myschool.edu/api/v1  # self-hosted only
canvas-cli auth set-token YOUR_TOKEN_HERE
canvas-cli doctor  # verify
```

## Command Reference

**account-calendars** — Manage account calendars

- `canvas-cli account-calendars api-index` — Returns a paginated list of account calendars available to the current user. Includes visible account calendars...
- `canvas-cli account-calendars api-show` — Get details about a specific account calendar.
- `canvas-cli account-calendars api-update` — Set an account calendar's visibility and auto_subscribe values. Requires the `manage_account_calendar_visibility`...

**accounts** — API for accessing account data.

- `canvas-cli accounts index` — A paginated list of accounts that the current user can view or manage. Typically, students and even teachers will...
- `canvas-cli accounts show` — Retrieve information on an individual account, given by id or sis sis_account_id.
- `canvas-cli accounts update` — Update an existing account.

**announcements** — Copyright (C) 2011 - present Instructure, Inc.

This file is part of Canvas.

Canvas is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, version 3 of the License.

Canvas is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
A PARTICULAR PURPOSE. See the GNU Affero General Public License for more
details.

You should have received a copy of the GNU Affero General Public License along
with this program. If not, see <http://www.gnu.org/licenses/>.

- `canvas-cli announcements` — Returns the paginated list of announcements for the given courses and date range. Note that a +context_code+ field...

**appointment-groups** — API for creating, accessing and updating appointment groups. Appointment groups
provide a way of creating a bundle of time slots that users can sign up for
(e.g. "Office Hours" or "Meet with professor about Final Project"). Both time
slots and reservations of time slots are stored as Calendar Events.

- `canvas-cli appointment-groups create` — Create and return a new appointment group. If new_appointments are specified, the response will return a...
- `canvas-cli appointment-groups destroy` — Delete an appointment group (and associated time slots and reservations) and return the deleted group
- `canvas-cli appointment-groups index` — Retrieve the paginated list of appointment groups that can be reserved or managed by the current user.
- `canvas-cli appointment-groups next-appointment` — Return the next appointment available to sign up for. The appointment is returned in a one-element array. If no...
- `canvas-cli appointment-groups show` — Returns information for a single appointment group
- `canvas-cli appointment-groups update` — Update and return an appointment group. If new_appointments are specified, the response will return a...

**audit** — Manage audit

- `canvas-cli audit authentication-api-for-account` — List authentication events for a given account.
- `canvas-cli audit authentication-api-for-login` — List authentication events for a given login.
- `canvas-cli audit authentication-api-for-user` — List authentication events for a given user.
- `canvas-cli audit course-api-for-account` — List course change events for a given account.
- `canvas-cli audit course-api-for-course` — List course change events for a given course.
- `canvas-cli audit grade-change-api-for-assignment` — List grade change events for a given assignment.
- `canvas-cli audit grade-change-api-for-course` — List grade change events for a given course.
- `canvas-cli audit grade-change-api-for-grader` — List grade change events for a given grader.
- `canvas-cli audit grade-change-api-for-student` — List grade change events for a given student.
- `canvas-cli audit grade-change-api-query` — List grade change events satisfying all given parameters. Teachers may query for events in courses they teach....

**brand-variables** — Manage brand variables

- `canvas-cli brand-variables` — Will redirect to a static json file that has all of the brand variables used by this account. Even though this is a...

**calendar-events** — Copyright (C) 2011 - present Instructure, Inc.

This file is part of Canvas.

Canvas is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, version 3 of the License.

Canvas is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
A PARTICULAR PURPOSE. See the GNU Affero General Public License for more
details.

You should have received a copy of the GNU Affero General Public License along
with this program. If not, see <http://www.gnu.org/licenses/>.

- `canvas-cli calendar-events api-create` — Create and return a new calendar event
- `canvas-cli calendar-events api-destroy` — Delete an event from the calendar and return the deleted event
- `canvas-cli calendar-events api-index` — Retrieve the paginated list of calendar events or assignments for the current user
- `canvas-cli calendar-events api-save-enabled-account-calendars` — Creates and updates the enabled_account_calendars and mark_feature_as_seen user preferences
- `canvas-cli calendar-events api-show` — Returns detailed information about a specific calendar event or assignment.
- `canvas-cli calendar-events api-update` — Update and return a calendar event

**canvadoc-session** — Copyright (C) 2014 - present Instructure, Inc.

This file is part of Canvas.

Canvas is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, version 3 of the License.

Canvas is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
A PARTICULAR PURPOSE. See the GNU Affero General Public License for more
details.

You should have received a copy of the GNU Affero General Public License along
with this program. If not, see <http://www.gnu.org/licenses/>.

- `canvas-cli canvadoc-session` — This API can only be accessed when another endpoint provides a signed URL. It will simply redirect you to the 3rd...

**canvas-lms-search** — Manage canvas lms search

- `canvas-cli canvas-lms-search all-courses` — A paginated list of all courses visible in the public index
- `canvas-cli canvas-lms-search recipients-other-2` — Find valid recipients (users, courses and groups) that the current user can send messages to. The...

**career** — Copyright (C) 2025 - present Instructure, Inc.

This file is part of Canvas.

Canvas is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, version 3 of the License.

Canvas is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
A PARTICULAR PURPOSE. See the GNU Affero General Public License for more
details.

You should have received a copy of the GNU Affero General Public License along
with this program. If not, see <http://www.gnu.org/licenses/>.

- `canvas-cli career experience-enabled` — Returns whether the root account has Canvas Career (Horizon) enabled in at least one subaccount.
- `canvas-cli career experience-experience-summary` — Returns the current user's active experience and available experiences they can switch to.
- `canvas-cli career experience-switch-experience` — Switch the current user's active experience to the specified one.
- `canvas-cli career experience-switch-role` — Switch the current user's role within the current experience.

**collaborations** — Manage collaborations


**comm-messages** — Manage comm messages

- `canvas-cli comm-messages` — Retrieve a paginated list of messages sent to a user.

**conferences** — API for accessing information on conferences.

- `canvas-cli conferences` — Retrieve the paginated list of conferences for all courses and groups the current user belongs to This API returns a...

**conversations** — API for creating, accessing and updating user conversations.

- `canvas-cli conversations batch-update` — Perform a change on a set of conversations. Operates asynchronously; use the {api:ProgressController#show progress...
- `canvas-cli conversations batches` — Returns any currently running conversation batches for the current user. Conversation batches are created when a...
- `canvas-cli conversations create` — Create a new conversation with one or more recipients. If there is already an existing private conversation with the...
- `canvas-cli conversations destroy` — Delete this conversation and its messages. Note that this only deletes this user's view of the conversation....
- `canvas-cli conversations index` — Returns the paginated list of conversations for the current user, most recent ones first....
- `canvas-cli conversations mark-all-as-read` — Mark all conversations as read.
- `canvas-cli conversations search-recipients-other` — Find valid recipients (users, courses and groups) that the current user can send messages to. The...
- `canvas-cli conversations show` — Returns information for a single conversation for the current user. Response includes all fields that are present in...
- `canvas-cli conversations unread-count` — Get the number of unread conversations for the current user
- `canvas-cli conversations update` — Updates attributes for a single conversation.

**course-accounts** — Manage course accounts

- `canvas-cli course-accounts` — A paginated list of accounts that the current user can view through their admin course enrollments. (Teacher, TA, or...

**course-creation-accounts** — Manage course creation accounts

- `canvas-cli course-creation-accounts` — A paginated list of accounts where the current user has permission to create courses.

**courses** — API for accessing course information.

- `canvas-cli courses destroy` — Delete or conclude an existing course
- `canvas-cli courses index` — Returns the paginated list of active courses for the current user.
- `canvas-cli courses show-other` — Return information on a single course. Accepts the same include[] parameters as the list action plus:
- `canvas-cli courses update` — Update an existing course. Arguments are the same as Courses#create, with a few exceptions (enroll_me). If a user...

**developer-keys** — Manage Canvas API Keys, used for OAuth access to this API.
See <a href="file.oauth.html">the OAuth access docs</a> for usage of these keys.
Note that DeveloperKeys are also (currently) used for LTI 1.3 registration and OIDC access,
but this endpoint deals with Canvas API keys. See <a href="file.registration.html">LTI Registration</a>
for details.

- `canvas-cli developer-keys destroy` — Delete an existing Canvas API key. Deleting an LTI 1.3 registration should be done via the LTI Registration API.
- `canvas-cli developer-keys update` — Update an existing Canvas API key. Updating an LTI 1.3 registration is not supported here and should be done via the...

**discovery-pages** — Manage discovery pages

- `canvas-cli discovery-pages api-show` — Get the discovery page configuration for the domain root account.
- `canvas-cli discovery-pages api-upsert` — Update or create the discovery page configuration for the domain root account. This is a full replacement - provide...

**enqueue-outcome-rollup-calculation** — Manage enqueue outcome rollup calculation

- `canvas-cli enqueue-outcome-rollup-calculation` — Enqueue a delayed Outcome Rollup Calculation Job

**eportfolios** — Copyright (C) 2011 - present Instructure, Inc.

This file is part of Canvas.

Canvas is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, version 3 of the License.

Canvas is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
A PARTICULAR PURPOSE. See the GNU Affero General Public License for more
details.

You should have received a copy of the GNU Affero General Public License along
with this program. If not, see <http://www.gnu.org/licenses/>.

- `canvas-cli eportfolios api-delete` — Mark an ePortfolio as deleted.
- `canvas-cli eportfolios api-show` — Get details for a single ePortfolio.

**epub-exports** — API for exporting courses as an ePub

- `canvas-cli epub-exports` — A paginated list of all courses a user is actively participating in, and the latest ePub export associated with the...

**error-reports** — Manage error reports

- `canvas-cli error-reports` — Create a new error report documenting an experienced problem Performs the same action as when a user uses the 'help...

**external-tools** — Manage external tools

- `canvas-cli external-tools` — Get a list of external tools with the course_navigation placement that have not been hidden in course settings and...

**features** — Manage features

- `canvas-cli features` — Return a hash of global feature options that pertain to the Canvas user interface. This is the same information...

**files** — Manage files

- `canvas-cli files metadata-sax-doc-api-show-other` — Returns the standard attachment json object
- `canvas-cli files metadata-sax-doc-api-update` — Update some settings on the specified file
- `canvas-cli files metadata-sax-doc-destroy` — Remove the specified file. Unlike most other DELETE endpoints, using this endpoint will result in comprehensive,...

**folders** — Manage folders

- `canvas-cli folders api-destroy` — Remove the specified folder. You can only delete empty folders unless you set the 'force' flag
- `canvas-cli folders show-other` — Returns the details for a folder You can get the root folder from a context by using 'root' as the :id. For example,...
- `canvas-cli folders update` — Updates a folder

**global** — Manage global

- `canvas-cli global outcome-groups-api-create-other` — Creates a new empty subgroup under the outcome group with the given title and description.
- `canvas-cli global outcome-groups-api-destroy-other` — Deleting an outcome group deletes descendant outcome groups and outcome links. The linked outcomes themselves are...
- `canvas-cli global outcome-groups-api-import-other` — Creates a new subgroup of the outcome group with the same title and description as the source group, then creates...
- `canvas-cli global outcome-groups-api-link-other` — Link an outcome into the outcome group. The outcome to link can either be specified by a PUT to the link URL for a...
- `canvas-cli global outcome-groups-api-link-other-2` — Link an outcome into the outcome group. The outcome to link can either be specified by a PUT to the link URL for a...
- `canvas-cli global outcome-groups-api-outcomes-other` — A paginated list of the immediate OutcomeLink children of the outcome group.
- `canvas-cli global outcome-groups-api-redirect-other` — Convenience redirect to find the root outcome group for a particular context. Will redirect to the appropriate...
- `canvas-cli global outcome-groups-api-show-other` — Returns detailed information about a specific outcome group.
- `canvas-cli global outcome-groups-api-subgroups-other` — A paginated list of the immediate OutcomeGroup children of the outcome group.
- `canvas-cli global outcome-groups-api-unlink-other` — Unlinking an outcome only deletes the outcome itself if this was the last link to the outcome in any group in any...
- `canvas-cli global outcome-groups-api-update-other` — Modify an existing outcome group. Fields not provided are left as is; unrecognized fields are ignored. When changing...

**grading-period-sets** — Manage grading period sets


**group-categories** — Group Categories allow grouping of groups together in canvas. There are a few
different built-in group categories used, or custom ones can be created. The
built in group categories are:  "communities", "student_organized", and "imported".

- `canvas-cli group-categories destroy` — Deletes a group category and all groups under it. Protected group categories can not be deleted, i.e. 'communities'...
- `canvas-cli group-categories show` — Returns the data for a single group category, or a 401 if the caller doesn't have the rights to see it.
- `canvas-cli group-categories update` — Modifies an existing group category.

**groups** — Groups serve as the data for a few different ideas in Canvas.  The first is
that they can be a community in the canvas network.  The second is that they
can be organized by students in a course, for study or communication (but not
grading).  The third is that they can be organized by teachers or account
administrators for the purpose of projects, assignments, and grading.  This
last kind of group is always part of a group category, which adds the
restriction that a user may only be a member of one group per category.

All of these types of groups function similarly, and can be the parent
context for many other types of functionality and interaction, such as
collections, discussions, wikis, and shared files.

- `canvas-cli groups create-other-2` — Creates a new group. Groups created using the '/api/v1/groups/' endpoint will be community groups.
- `canvas-cli groups destroy` — Deletes a group and removes all members.
- `canvas-cli groups show` — Returns the data for a single group, or a 401 if the caller doesn't have the rights to see it.
- `canvas-cli groups update` — Modifies an existing group. Note that to set an avatar image for the group, you must first upload the image file to...

**inst-access-tokens** — Short term JWT tokens that can be used to authenticate with Canvas and other
Instructure services.  InstAccess tokens expire after one hour.  Canvas hands
out encrypted tokens that need to be decrypted by the API Gateway before they
can be accepted by Canvas or other services.

- `canvas-cli inst-access-tokens` — Create a unique, encrypted InstAccess token. Generates a different InstAccess token each time it's called, each one...

**jwts** — Short term tokens useful for talking to other services in the Canvas Ecosystem.
Note: JWTs have no value or use directly against the Canvas API, and expire
after one hour

- `canvas-cli jwts create` — Create a unique JWT for use with other Canvas services Generates a different JWT each time it's called. Each JWT...
- `canvas-cli jwts refresh` — Refresh a JWT for use with other canvas services Generates a different JWT each time it's called, each one expires...

**manageable-accounts** — Manage manageable accounts

- `canvas-cli manageable-accounts` — A paginated list of accounts where the current user has permission to create or manage courses. List will be empty...

**manually-created-courses-account** — Manage manually created courses account

- `canvas-cli manually-created-courses-account` — Returns the sub-account that contains manually created courses for the domain root account.

**media-attachments** — Manage media attachments

- `canvas-cli media-attachments media-objects-index-other-2` — Returns media objects created by the user making the request. When using the second version, returns media objects...
- `canvas-cli media-attachments media-objects-update-media-object-other-2` — Updates the title of a media object.

**media-objects** — When you upload or record webcam video/audio to kaltura, it makes a Media Object

- `canvas-cli media-objects index-other` — Returns media objects created by the user making the request. When using the second version, returns media objects...
- `canvas-cli media-objects update-other` — Updates the title of a media object.

**outcomes** — Copyright (C) 2011 - present Instructure, Inc.

This file is part of Canvas.

Canvas is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, version 3 of the License.

Canvas is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
A PARTICULAR PURPOSE. See the GNU Affero General Public License for more
details.

You should have received a copy of the GNU Affero General Public License along
with this program. If not, see <http://www.gnu.org/licenses/>.

- `canvas-cli outcomes api-show` — Returns the details of the outcome with the given id.
- `canvas-cli outcomes api-update` — Modify an existing outcome. Fields not provided are left as is; unrecognized fields are ignored. If any new ratings...

**permissions** — Manage permissions

- `canvas-cli permissions` — Retrieve information about groups of granular permissions The return value is a dictionary of permission group keys...

**planner** — API for listing learning objects to display on the student planner and calendar

- `canvas-cli planner index-other` — Retrieve the paginated list of objects to be shown on the planner for the current user with the associated planner...
- `canvas-cli planner overrides-create` — Create a planner override for the current user
- `canvas-cli planner overrides-destroy` — Delete a planner override for the current user
- `canvas-cli planner overrides-index` — Retrieve a planner override for the current user
- `canvas-cli planner overrides-show` — Retrieve a planner override for the current user
- `canvas-cli planner overrides-update` — Update a planner override's visibilty for the current user

**planner-notes** — API for creating, accessing and updating Planner Notes. PlannerNote are used
to set reminders and notes to self about courses or general events.

- `canvas-cli planner-notes create` — Create a planner note for the current user
- `canvas-cli planner-notes destroy` — Delete a planner note for the current user
- `canvas-cli planner-notes index` — Retrieve the paginated list of planner notes Retrieve planner note for a user
- `canvas-cli planner-notes show` — Retrieve a planner note for the current user
- `canvas-cli planner-notes update` — Update a planner note for the current user

**progress** — API for querying the progress of asynchronous API operations.

- `canvas-cli progress <id>` — Return completion and status information about an asynchronous job

**question-banks** — Copyright (C) 2011 - present Instructure, Inc.

This file is part of Canvas.

Canvas is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, version 3 of the License.

Canvas is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
A PARTICULAR PURPOSE. See the GNU Affero General Public License for more
details.

You should have received a copy of the GNU Affero General Public License along
with this program. If not, see <http://www.gnu.org/licenses/>.

- `canvas-cli question-banks assessment-index` — Returns the paginated list of question banks for a given context.
- `canvas-cli question-banks assessment-show` — Returns the question bank with the given id

**rich-content** — Manage rich content

- `canvas-cli rich-content` — Generates a rich content.

**rubrics** — Manage rubrics

- `canvas-cli rubrics` — Returns a CSV template file that can be used to import rubrics into Canvas.

**sections** — API for accessing section information.

- `canvas-cli sections destroy` — Delete an existing section. Returns the former Section.
- `canvas-cli sections show-other` — Gets details about a specific section
- `canvas-cli sections update` — Modify an existing section.

**services** — Manage services

- `canvas-cli services api-show-kaltura-config` — Return the config information for the Kaltura plugin in json format.
- `canvas-cli services api-start-kaltura-session` — Start a new Kaltura session, so that new media can be recorded and uploaded to this Canvas instance's Kaltura instance.

**settings** — Manage settings

- `canvas-cli settings` — Return a hash of global settings for the root account This is the same information supplied to the web interface as...

**shared-brand-configs** — This is how you can share Themes with other people in your account or
so you can come back to them later without having to apply them to your account

- `canvas-cli shared-brand-configs <id>` — Delete a SharedBrandConfig, which will unshare it so you nor anyone else in your account will see it as an option to...

**submissions** — Manage submissions


**users** — Manage users

- `canvas-cli users communication-channels-delete-push-token` — Delete a push notification endpoint
- `canvas-cli users course-nicknames-clear` — Remove all stored course nicknames.
- `canvas-cli users course-nicknames-delete` — Remove the nickname for the given course. Subsequent course API calls will return the actual name for the course.
- `canvas-cli users course-nicknames-index` — Returns all course nicknames you have set.
- `canvas-cli users course-nicknames-show` — Returns the nickname for a specific course.
- `canvas-cli users course-nicknames-update` — Set a nickname for the given course. This will replace the course's name in output of API calls you make...
- `canvas-cli users favorites-add-favorite-course` — Add a course to the current user's favorites. If the course is already in the user's favorites, nothing happens....
- `canvas-cli users favorites-add-favorite-groups` — Add a group to the current user's favorites. If the group is already in the user's favorites, nothing happens.
- `canvas-cli users favorites-list-favorite-courses` — Retrieve the paginated list of favorite courses for the current user. If the user has not chosen any favorites, then...
- `canvas-cli users favorites-list-favorite-groups` — Retrieve the paginated list of favorite groups for the current user. If the user has not chosen any favorites, then...
- `canvas-cli users favorites-remove-favorite-course` — Remove a course from the current user's favorites.
- `canvas-cli users favorites-remove-favorite-groups` — Remove a group from the current user's favorites.
- `canvas-cli users favorites-reset-course-favorites` — Reset the current user's course favorites to the default automatically generated list of enrolled courses
- `canvas-cli users favorites-reset-groups-favorites` — Reset the current user's group favorites to the default automatically generated list of enrolled group
- `canvas-cli users groups-index` — Returns a paginated list of active groups for the current user.
- `canvas-cli users notification-preferences-update-all-for` — Change the preferences for multiple notifications for a single communication channel at once
- `canvas-cli users notification-preferences-update-all-for-2` — Change the preferences for multiple notifications for a single communication channel at once
- `canvas-cli users notification-preferences-update-for` — Change the preference for a single notification for a single communication channel
- `canvas-cli users notification-preferences-update-for-2` — Change the preference for a single notification for a single communication channel
- `canvas-cli users notification-preferences-update-preferences-by-category` — Change the preferences for multiple notifications based on the category for a single communication channel
- `canvas-cli users page-views-batch-query` — Initiates an asynchronous query for page views data across multiple users. This method enqueues a background job to...
- `canvas-cli users page-views-batch-query-results` — Retrieves the results of a completed batch page views query. Returns the data in the format specified when the query...
- `canvas-cli users page-views-poll-batch-query` — Checks the status of a previously initiated batch page views query. Returns the current processing status and...
- `canvas-cli users pseudonyms-forgot-password` — Given a user email, generate a nonce and email it to the user
- `canvas-cli users service-credentials-activity-stream-for` — Returns the current user's global activity stream, paginated. There are many types of objects that can be returned...
- `canvas-cli users service-credentials-activity-stream-other` — Returns the current user's global activity stream, paginated. There are many types of objects that can be returned...
- `canvas-cli users service-credentials-activity-stream-summary` — Returns a summary of the current user's global activity stream.
- `canvas-cli users service-credentials-api-show` — Shows details for user. Also includes an attribute 'permissions', a non-comprehensive list of permissions for the...
- `canvas-cli users service-credentials-expire-mobile-sessions-other` — Permanently expires any active mobile sessions, forcing them to re-authorize. The route that takes a user id will...
- `canvas-cli users service-credentials-ignore-all-stream-items` — Hide all stream items for the user
- `canvas-cli users service-credentials-ignore-stream-item` — Hide the given stream item.
- `canvas-cli users service-credentials-pandata-events-token` — Returns a jwt auth and props token that can be used to send events to Pandata. NOTE: This is currently only...
- `canvas-cli users service-credentials-todo-item-count` — Counts of different todo items such as the number of assignments needing grading as well as the number of...
- `canvas-cli users service-credentials-todo-items` — A paginated list of the current user's list of todo items. There is a limit to the number of items returned. The...
- `canvas-cli users service-credentials-upcoming-events` — A paginated list of the current user's upcoming events.
- `canvas-cli users service-credentials-update` — Modify an existing user. To modify a user's login, see the documentation for logins.


### Finding the right command

When you know what you want to do but not which command does it, ask the CLI directly:

```bash
canvas-cli which "<capability in your own words>"
```

`which` resolves a natural-language capability query to the best matching command from this CLI's curated feature index. Exit code `0` means at least one match; exit code `2` means no confident match — fall back to `--help` or use a narrower query.

## Auth Setup

Store your access token:

```bash
canvas-cli auth set-token YOUR_TOKEN_HERE
```

Or set `CANVAS_LMS_TOKEN` as an environment variable.

Run `canvas-cli doctor` to verify setup.

## Agent Mode

Add `--agent` to any command. Expands to: `--json --compact --no-input --no-color --yes`.

- **Pipeable** — JSON on stdout, errors on stderr
- **Filterable** — `--select` keeps a subset of fields. Dotted paths descend into nested structures; arrays traverse element-wise. Critical for keeping context small on verbose APIs:

  ```bash
  canvas-cli account-calendars api-index --agent --select id,name,status
  ```
- **Previewable** — `--dry-run` shows the request without sending
- **Offline-friendly** — sync/search commands can use the local SQLite store when available
- **Non-interactive** — never prompts, every input is a flag
- **Explicit retries** — use `--idempotent` only when an already-existing create should count as success, and `--ignore-missing` only when a missing delete target should count as success

### Response envelope

Commands that read from the local store or the API wrap output in a provenance envelope:

```json
{
  "meta": {"source": "live" | "local", "synced_at": "...", "reason": "..."},
  "results": <data>
}
```

Parse `.results` for data and `.meta.source` to know whether it's live or local. A human-readable `N results (live)` summary is printed to stderr only when stdout is a terminal — piped/agent consumers get pure JSON on stdout.

## Agent Feedback

When you (or the agent) notice something off about this CLI, record it:

```
canvas-cli feedback "the --since flag is inclusive but docs say exclusive"
canvas-cli feedback --stdin < notes.txt
canvas-cli feedback list --json --limit 10
```

Entries are stored locally at `~/.canvas-cli/feedback.jsonl`. They are never POSTed unless `CANVAS_FEEDBACK_ENDPOINT` is set AND either `--send` is passed or `CANVAS_FEEDBACK_AUTO_SEND=true`. Default behavior is local-only.

Write what *surprised* you, not a bug report. Short, specific, one line: that is the part that compounds.

## Output Delivery

Every command accepts `--deliver <sink>`. The output goes to the named sink in addition to (or instead of) stdout, so agents can route command results without hand-piping. Three sinks are supported:

| Sink | Effect |
|------|--------|
| `stdout` | Default; write to stdout only |
| `file:<path>` | Atomically write output to `<path>` (tmp + rename) |
| `webhook:<url>` | POST the output body to the URL (`application/json` or `application/x-ndjson` when `--compact`) |

Unknown schemes are refused with a structured error naming the supported set. Webhook failures return non-zero and log the URL + HTTP status on stderr.

## Named Profiles

A profile is a saved set of flag values, reused across invocations. Use it when a scheduled agent calls the same command every run with the same configuration - HeyGen's "Beacon" pattern.

```
canvas-cli profile save briefing --json
canvas-cli --profile briefing account-calendars api-index
canvas-cli profile list --json
canvas-cli profile show briefing
canvas-cli profile delete briefing --yes
```

Explicit flags always win over profile values; profile values win over defaults. `agent-context` lists all available profiles under `available_profiles` so introspecting agents discover them at runtime.

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 2 | Usage error (wrong arguments) |
| 3 | Resource not found |
| 4 | Authentication required |
| 5 | API error (upstream issue) |
| 7 | Rate limited (wait and retry) |
| 10 | Config error |

## Argument Parsing

Parse `$ARGUMENTS`:

1. **Empty, `help`, or `--help`** → show `canvas-cli --help` output
2. **Starts with `install`** → ends with `mcp` → MCP installation; otherwise → see Prerequisites above
3. **Anything else** → Direct Use (execute as CLI command with `--agent`)

## MCP Server Installation

Install the MCP binary from this CLI's published public-library entry or pre-built release, then register it:

```bash
claude mcp add canvas-mcp -- canvas-mcp
```

Verify: `claude mcp list`

## Direct Use

1. Check if installed: `which canvas-cli`
   If not found, offer to install (see Prerequisites at the top of this skill).
2. Match the user query to the best command from the Unique Capabilities and Command Reference above.
3. Execute with the `--agent` flag:
   ```bash
   canvas-cli <command> [subcommand] [args] --agent
   ```
4. If ambiguous, drill into subcommand help: `canvas-cli <command> --help`.
