# Canvas CLI

A full-featured CLI and MCP server for the Canvas LMS API, with built-in commands for every Canvas resource plus five compound intelligence commands.

## Install

```bash
curl -sSfL https://raw.githubusercontent.com/simple-scalable-solutions/canvas-cli/main/install.sh | sh
```

This installs both `canvas-cli` and `canvas-mcp` to `/usr/local/bin`. To install to a different directory:

```bash
INSTALL_DIR=~/.local/bin curl -sSfL https://raw.githubusercontent.com/simple-scalable-solutions/canvas-cli/main/install.sh | sh
```

## Install Claude Code Skill

To let Claude Code use canvas-cli on your behalf, install the skill:

```bash
curl -sSfL https://raw.githubusercontent.com/simple-scalable-solutions/canvas-cli/main/install-skill.sh | sh
```

Then restart Claude Code. Claude will automatically use `canvas-cli` when you ask it to interact with Canvas.

## Quick Start

### 1. Save your Canvas URL (self-hosted only)

Skip this step if you use Canvas Cloud (`*.instructure.com`).

```bash
canvas-cli auth set-url https://canvas.myschool.edu/api/v1
```

### 2. Save your API token

Get your token from Canvas: **Account → Settings → Approved Integrations → New Access Token**

```bash
canvas-cli auth set-token YOUR_TOKEN_HERE
```

### 3. Verify setup

```bash
canvas-cli auth status
canvas-cli doctor
```

### 4. Try your first command

```bash
canvas-cli account-calendars api-index
```

## Usage

Run `canvas-cli --help` for the full command reference and flag list.

## Commands

### account-calendars

Manage account calendars

- **`canvas-cli account-calendars api-index`** - Returns a paginated list of account calendars available to the current user.
Includes visible account calendars where the user has an account association.
- **`canvas-cli account-calendars api-show`** - Get details about a specific account calendar.
- **`canvas-cli account-calendars api-update`** - Set an account calendar's visibility and auto_subscribe values. Requires the
`manage_account_calendar_visibility` permission on the account.

### accounts

API for accessing account data.

- **`canvas-cli accounts index`** - A paginated list of accounts that the current user can view or manage.
Typically, students and even teachers will get an empty list in response,
only account admins can view the accounts that they are in.
- **`canvas-cli accounts show`** - Retrieve information on an individual account, given by id or sis
sis_account_id.
- **`canvas-cli accounts update`** - Update an existing account.

### announcements

Copyright (C) 2011 - present Instructure, Inc.

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

- **`canvas-cli announcements api-index`** - Returns the paginated list of announcements for the given courses and date range.  Note that
a +context_code+ field is added to the responses so you can tell which course each announcement
belongs to.

### appointment-groups

API for creating, accessing and updating appointment groups. Appointment groups
provide a way of creating a bundle of time slots that users can sign up for
(e.g. "Office Hours" or "Meet with professor about Final Project"). Both time
slots and reservations of time slots are stored as Calendar Events.

- **`canvas-cli appointment-groups create`** - Create and return a new appointment group. If new_appointments are
specified, the response will return a new_appointments array (same format
as appointments array, see "List appointment groups" action)
- **`canvas-cli appointment-groups destroy`** - Delete an appointment group (and associated time slots and reservations)
and return the deleted group
- **`canvas-cli appointment-groups index`** - Retrieve the paginated list of appointment groups that can be reserved or
managed by the current user.
- **`canvas-cli appointment-groups next-appointment`** - Return the next appointment available to sign up for. The appointment
is returned in a one-element array. If no future appointments are
available, an empty array is returned.
- **`canvas-cli appointment-groups show`** - Returns information for a single appointment group
- **`canvas-cli appointment-groups update`** - Update and return an appointment group. If new_appointments are specified,
the response will return a new_appointments array (same format as
appointments array, see "List appointment groups" action).

### audit

Manage audit

- **`canvas-cli audit authentication-api-for-account`** - List authentication events for a given account.
- **`canvas-cli audit authentication-api-for-login`** - List authentication events for a given login.
- **`canvas-cli audit authentication-api-for-user`** - List authentication events for a given user.
- **`canvas-cli audit course-api-for-account`** - List course change events for a given account.
- **`canvas-cli audit course-api-for-course`** - List course change events for a given course.
- **`canvas-cli audit grade-change-api-for-assignment`** - List grade change events for a given assignment.
- **`canvas-cli audit grade-change-api-for-course`** - List grade change events for a given course.
- **`canvas-cli audit grade-change-api-for-grader`** - List grade change events for a given grader.
- **`canvas-cli audit grade-change-api-for-student`** - List grade change events for a given student.
- **`canvas-cli audit grade-change-api-query`** - List grade change events satisfying all given parameters. Teachers may query for events in courses they teach.
Queries without +course_id+ require account administrator rights.

At least one of +course_id+, +assignment_id+, +student_id+, or +grader_id+ must be specified.

### brand-variables

Manage brand variables

- **`canvas-cli brand-variables brand-configs-api-show`** - Will redirect to a static json file that has all of the brand
variables used by this account. Even though this is a redirect,
do not store the redirected url since if the account makes any changes
it will redirect to a new url. Needs no authentication.

### calendar-events

Copyright (C) 2011 - present Instructure, Inc.

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

- **`canvas-cli calendar-events api-create`** - Create and return a new calendar event
- **`canvas-cli calendar-events api-destroy`** - Delete an event from the calendar and return the deleted event
- **`canvas-cli calendar-events api-index`** - Retrieve the paginated list of calendar events or assignments for the current user
- **`canvas-cli calendar-events api-save-enabled-account-calendars`** - Creates and updates the enabled_account_calendars and mark_feature_as_seen user preferences
- **`canvas-cli calendar-events api-show`** - Returns detailed information about a specific calendar event or assignment.
- **`canvas-cli calendar-events api-update`** - Update and return a calendar event

### canvadoc-session

Copyright (C) 2014 - present Instructure, Inc.

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

- **`canvas-cli canvadoc-session show`** - This API can only be accessed when another endpoint provides a signed URL.
It will simply redirect you to the 3rd party document preview.

### canvas-lms-search

Manage canvas lms search

- **`canvas-cli canvas-lms-search all-courses`** - A paginated list of all courses visible in the public index
- **`canvas-cli canvas-lms-search recipients-other-2`** - Find valid recipients (users, courses and groups) that the current user
can send messages to. The /api/v1/search/recipients path is the preferred
endpoint, /api/v1/conversations/find_recipients is deprecated.

Pagination is supported.

### career

Copyright (C) 2025 - present Instructure, Inc.

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

- **`canvas-cli career experience-enabled`** - Returns whether the root account has Canvas Career (Horizon) enabled
in at least one subaccount.
- **`canvas-cli career experience-experience-summary`** - Returns the current user's active experience and available experiences
they can switch to.
- **`canvas-cli career experience-switch-experience`** - Switch the current user's active experience to the specified one.
- **`canvas-cli career experience-switch-role`** - Switch the current user's role within the current experience.

### collaborations

Manage collaborations


### comm-messages

Manage comm messages

- **`canvas-cli comm-messages api-index`** - Retrieve a paginated list of messages sent to a user.

### conferences

API for accessing information on conferences.

- **`canvas-cli conferences for-user`** - Retrieve the paginated list of conferences for all courses and groups
the current user belongs to

This API returns a JSON object containing the list of conferences.
The key for the list of conferences is "conferences".

### conversations

API for creating, accessing and updating user conversations.

- **`canvas-cli conversations batch-update`** - Perform a change on a set of conversations. Operates asynchronously; use the {api:ProgressController#show progress endpoint}
to query the status of an operation.
- **`canvas-cli conversations batches`** - Returns any currently running conversation batches for the current user.
Conversation batches are created when a bulk private message is sent
asynchronously (see the mode argument to the {api:ConversationsController#create create API action}).
- **`canvas-cli conversations create`** - Create a new conversation with one or more recipients. If there is already
an existing private conversation with the given recipients, it will be
reused.

 (either numeric IDs or UUIDs prefixed with "uuid:"),
  or course/group ids prefixed with "course_" or "group_" respectively, e.g.
  recipients[]=1&recipients[]=uuid:W9GQIcdoDTqwX8mxIunDQQVL6WZTaGmpa5xovmCBx&recipients[]=course_3.
  If the course/group has over 100 enrollments, 'bulk_message' and 'group_conversation' must be
  set to true.
- **`canvas-cli conversations destroy`** - Delete this conversation and its messages. Note that this only deletes
this user's view of the conversation.

Response includes same fields as UPDATE action
- **`canvas-cli conversations index`** - Returns the paginated list of conversations for the current user, most
recent ones first.

 "uuid:W9GQIcdoDTqwX8mxIunDQQVL6WZTaGmpa5xovmCB", or "course_456".
 For users, you can use either their numeric ID or UUID prefixed with "uuid:".
 Can be an array (by setting "filter[]") or single value (by setting "filter")
- **`canvas-cli conversations mark-all-as-read`** - Mark all conversations as read.
- **`canvas-cli conversations search-recipients-other`** - Find valid recipients (users, courses and groups) that the current user
can send messages to. The /api/v1/search/recipients path is the preferred
endpoint, /api/v1/conversations/find_recipients is deprecated.

Pagination is supported.
- **`canvas-cli conversations show`** - Returns information for a single conversation for the current user. Response includes all
fields that are present in the list/index action as well as messages
and extended participant information.
- **`canvas-cli conversations unread-count`** - Get the number of unread conversations for the current user
- **`canvas-cli conversations update`** - Updates attributes for a single conversation.

### course-accounts

Manage course accounts

- **`canvas-cli course-accounts accounts`** - A paginated list of accounts that the current user can view through their
admin course enrollments. (Teacher, TA, or designer enrollments).
Only returns "id", "name", "workflow_state", "root_account_id" and "parent_account_id"

### course-creation-accounts

Manage course creation accounts

- **`canvas-cli course-creation-accounts accounts`** - A paginated list of accounts where the current user has permission to create
courses.

### courses

API for accessing course information.

- **`canvas-cli courses destroy`** - Delete or conclude an existing course
- **`canvas-cli courses index`** - Returns the paginated list of active courses for the current user.
- **`canvas-cli courses show-other`** - Return information on a single course.

Accepts the same include[] parameters as the list action plus:
- **`canvas-cli courses update`** - Update an existing course.

Arguments are the same as Courses#create, with a few exceptions (enroll_me).

If a user has content management rights, but not full course editing rights, the only attribute
editable through this endpoint will be "syllabus_body"

If an account has set prevent_course_availability_editing_by_teachers, a teacher cannot change
+course[start_at]+, +course[conclude_at]+, or +course[restrict_enrollments_to_course_dates]+ here.

### developer-keys

Manage Canvas API Keys, used for OAuth access to this API.
See <a href="file.oauth.html">the OAuth access docs</a> for usage of these keys.
Note that DeveloperKeys are also (currently) used for LTI 1.3 registration and OIDC access,
but this endpoint deals with Canvas API keys. See <a href="file.registration.html">LTI Registration</a>
for details.

- **`canvas-cli developer-keys destroy`** - Delete an existing Canvas API key. Deleting an LTI 1.3 registration should be done via the LTI Registration API.
- **`canvas-cli developer-keys update`** - Update an existing Canvas API key. Updating an LTI 1.3 registration is not supported here and should
be done via the LTI Registration API.

### discovery-pages

Manage discovery pages

- **`canvas-cli discovery-pages api-show`** - Get the discovery page configuration for the domain root account.
- **`canvas-cli discovery-pages api-upsert`** - Update or create the discovery page configuration for the domain root account.
This is a full replacement - provide the complete configuration including
primary, secondary, and active fields. Any fields omitted will be removed.

### enqueue-outcome-rollup-calculation

Manage enqueue outcome rollup calculation

- **`canvas-cli enqueue-outcome-rollup-calculation outcome-results`** - Enqueue a delayed Outcome Rollup Calculation Job

### eportfolios

Copyright (C) 2011 - present Instructure, Inc.

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

- **`canvas-cli eportfolios api-delete`** - Mark an ePortfolio as deleted.
- **`canvas-cli eportfolios api-show`** - Get details for a single ePortfolio.

### epub-exports

API for exporting courses as an ePub

- **`canvas-cli epub-exports index`** - A paginated list of all courses a user is actively participating in, and
the latest ePub export associated with the user & course.

### error-reports

Manage error reports

- **`canvas-cli error-reports errors-create`** - Create a new error report documenting an experienced problem

Performs the same action as when a user uses the "help -> report a problem"
dialog.

### external-tools

Manage external tools

- **`canvas-cli external-tools timing-meta-all-visible-nav-tools`** - Get a list of external tools with the course_navigation placement that have not been hidden in
course settings and whose visibility settings apply to the requesting user. These tools are the
same that appear in the course navigation.

The response format is the same as for List external tools, but with additional context_id and
context_name fields on each element in the array.

### features

Manage features

- **`canvas-cli features flags-environment`** - Return a hash of global feature options that pertain to the
Canvas user interface. This is the same information supplied to the
web interface as +ENV.FEATURES+.

### files

Manage files

- **`canvas-cli files metadata-sax-doc-api-show-other`** - Returns the standard attachment json object
- **`canvas-cli files metadata-sax-doc-api-update`** - Update some settings on the specified file
- **`canvas-cli files metadata-sax-doc-destroy`** - Remove the specified file. Unlike most other DELETE endpoints, using this
endpoint will result in comprehensive, irretrievable destruction of the file.
It should be used with the `replace` parameter set to true in cases where the
file preview also needs to be destroyed (such as to remove files that violate
privacy laws).

### folders

Manage folders

- **`canvas-cli folders api-destroy`** - Remove the specified folder. You can only delete empty folders unless you
set the 'force' flag
- **`canvas-cli folders show-other`** - Returns the details for a folder

You can get the root folder from a context by using 'root' as the :id.
For example, you could get the root folder for a course like:
- **`canvas-cli folders update`** - Updates a folder

### global

Manage global

- **`canvas-cli global outcome-groups-api-create-other`** - Creates a new empty subgroup under the outcome group with the given title
and description.
- **`canvas-cli global outcome-groups-api-destroy-other`** - Deleting an outcome group deletes descendant outcome groups and outcome
links. The linked outcomes themselves are only deleted if all links to the
outcome were deleted.

Aligned outcomes cannot be deleted; as such, if all remaining links to an
aligned outcome are included in this group's descendants, the group
deletion will fail.
- **`canvas-cli global outcome-groups-api-import-other`** - Creates a new subgroup of the outcome group with the same title and
description as the source group, then creates links in that new subgroup to
the same outcomes that are linked in the source group. Recurses on the
subgroups of the source group, importing them each in turn into the new
subgroup.

Allows you to copy organizational structure, but does not create copies of
the outcomes themselves, only new links.

The source group must be either global, from the same context as this
outcome group, or from an associated account. The source group cannot be
the root outcome group of its context.
- **`canvas-cli global outcome-groups-api-link-other`** - Link an outcome into the outcome group. The outcome to link can either be
specified by a PUT to the link URL for a specific outcome (the outcome_id
in the PUT URLs) or by supplying the information for a new outcome (title,
description, ratings, mastery_points) in a POST to the collection.

If linking an existing outcome, the outcome_id must identify an outcome
available to this context; i.e. an outcome owned by this group's context,
an outcome owned by an associated account, or a global outcome. With
outcome_id present, any other parameters (except move_from) are ignored.

If defining a new outcome, the outcome is created in the outcome group's
context using the provided title, description, ratings, and mastery points;
the title is required but all other fields are optional. The new outcome
is then linked into the outcome group.

If ratings are provided when creating a new outcome, an embedded rubric
criterion is included in the new outcome. This criterion's mastery_points
default to the maximum points in the highest rating if not specified in the
mastery_points parameter. Any ratings lacking a description are given a
default of "No description". Any ratings lacking a point value are given a
default of 0. If no ratings are provided, the mastery_points parameter is
ignored.
- **`canvas-cli global outcome-groups-api-link-other-2`** - Link an outcome into the outcome group. The outcome to link can either be
specified by a PUT to the link URL for a specific outcome (the outcome_id
in the PUT URLs) or by supplying the information for a new outcome (title,
description, ratings, mastery_points) in a POST to the collection.

If linking an existing outcome, the outcome_id must identify an outcome
available to this context; i.e. an outcome owned by this group's context,
an outcome owned by an associated account, or a global outcome. With
outcome_id present, any other parameters (except move_from) are ignored.

If defining a new outcome, the outcome is created in the outcome group's
context using the provided title, description, ratings, and mastery points;
the title is required but all other fields are optional. The new outcome
is then linked into the outcome group.

If ratings are provided when creating a new outcome, an embedded rubric
criterion is included in the new outcome. This criterion's mastery_points
default to the maximum points in the highest rating if not specified in the
mastery_points parameter. Any ratings lacking a description are given a
default of "No description". Any ratings lacking a point value are given a
default of 0. If no ratings are provided, the mastery_points parameter is
ignored.
- **`canvas-cli global outcome-groups-api-outcomes-other`** - A paginated list of the immediate OutcomeLink children of the outcome group.
- **`canvas-cli global outcome-groups-api-redirect-other`** - Convenience redirect to find the root outcome group for a particular
context. Will redirect to the appropriate outcome group's URL.
- **`canvas-cli global outcome-groups-api-show-other`** - Returns detailed information about a specific outcome group.
- **`canvas-cli global outcome-groups-api-subgroups-other`** - A paginated list of the immediate OutcomeGroup children of the outcome group.
- **`canvas-cli global outcome-groups-api-unlink-other`** - Unlinking an outcome only deletes the outcome itself if this was the last
link to the outcome in any group in any context. Aligned outcomes cannot be
deleted; as such, if this is the last link to an aligned outcome, the
unlinking will fail.
- **`canvas-cli global outcome-groups-api-update-other`** - Modify an existing outcome group. Fields not provided are left as is;
unrecognized fields are ignored.

When changing the parent outcome group, the new parent group must belong to
the same context as this outcome group, and must not be a descendant of
this outcome group (i.e. no cycles allowed).

### grading-period-sets

Manage grading period sets


### group-categories

Group Categories allow grouping of groups together in canvas. There are a few
different built-in group categories used, or custom ones can be created. The
built in group categories are:  "communities", "student_organized", and "imported".

- **`canvas-cli group-categories destroy`** - Deletes a group category and all groups under it. Protected group
categories can not be deleted, i.e. "communities" and "student_organized".
- **`canvas-cli group-categories show`** - Returns the data for a single group category, or a 401 if the caller doesn't have
the rights to see it.
- **`canvas-cli group-categories update`** - Modifies an existing group category.

### groups

Groups serve as the data for a few different ideas in Canvas.  The first is
that they can be a community in the canvas network.  The second is that they
can be organized by students in a course, for study or communication (but not
grading).  The third is that they can be organized by teachers or account
administrators for the purpose of projects, assignments, and grading.  This
last kind of group is always part of a group category, which adds the
restriction that a user may only be a member of one group per category.

All of these types of groups function similarly, and can be the parent
context for many other types of functionality and interaction, such as
collections, discussions, wikis, and shared files.

- **`canvas-cli groups create-other-2`** - Creates a new group. Groups created using the "/api/v1/groups/"
endpoint will be community groups.
- **`canvas-cli groups destroy`** - Deletes a group and removes all members.
- **`canvas-cli groups show`** - Returns the data for a single group, or a 401 if the caller doesn't have
the rights to see it.
- **`canvas-cli groups update`** - Modifies an existing group.  Note that to set an avatar image for the
group, you must first upload the image file to the group, and the use the
id in the response as the argument to this function.  See the
{file:file.file_uploads.html File Upload Documentation} for details on the file
upload workflow.

### inst-access-tokens

Short term JWT tokens that can be used to authenticate with Canvas and other
Instructure services.  InstAccess tokens expire after one hour.  Canvas hands
out encrypted tokens that need to be decrypted by the API Gateway before they
can be accepted by Canvas or other services.

- **`canvas-cli inst-access-tokens create`** - Create a unique, encrypted InstAccess token.

Generates a different InstAccess token each time it's called, each one expires
after a short window (1 hour).

### jwts

Short term tokens useful for talking to other services in the Canvas Ecosystem.
Note: JWTs have no value or use directly against the Canvas API, and expire
after one hour

- **`canvas-cli jwts create`** - Create a unique JWT for use with other Canvas services

Generates a different JWT each time it's called. Each JWT expires
after a short window (1 hour)
- **`canvas-cli jwts refresh`** - Refresh a JWT for use with other canvas services

Generates a different JWT each time it's called, each one expires
after a short window (1 hour).

### manageable-accounts

Manage manageable accounts

- **`canvas-cli manageable-accounts accounts`** - A paginated list of accounts where the current user has permission to create
or manage courses. List will be empty for students and teachers as only admins
can view which accounts they are in.

### manually-created-courses-account

Manage manually created courses account

- **`canvas-cli manually-created-courses-account accounts`** - Returns the sub-account that contains manually created courses for the domain root account.

### media-attachments

Manage media attachments

- **`canvas-cli media-attachments media-objects-index-other-2`** - Returns media objects created by the user making the request. When
using the second version, returns media objects associated with
the given course.
- **`canvas-cli media-attachments media-objects-update-media-object-other-2`** - Updates the title of a media object.

### media-objects

When you upload or record webcam video/audio to kaltura, it makes a Media Object

- **`canvas-cli media-objects index-other`** - Returns media objects created by the user making the request. When
using the second version, returns media objects associated with
the given course.
- **`canvas-cli media-objects update-other`** - Updates the title of a media object.

### outcomes

Copyright (C) 2011 - present Instructure, Inc.

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

- **`canvas-cli outcomes api-show`** - Returns the details of the outcome with the given id.
- **`canvas-cli outcomes api-update`** - Modify an existing outcome. Fields not provided are left as is;
unrecognized fields are ignored.

If any new ratings are provided, the combination of all new ratings
provided completely replace any existing embedded rubric criterion; it is
not possible to tweak the ratings of the embedded rubric criterion.

A new embedded rubric criterion's mastery_points default to the maximum
points in the highest rating if not specified in the mastery_points
parameter. Any new ratings lacking a description are given a default of "No
description". Any new ratings lacking a point value are given a default of
0.

### permissions

Manage permissions

- **`canvas-cli permissions help-groups`** - Retrieve information about groups of granular permissions

The return value is a dictionary of permission group keys to objects
containing +label+ and +subtitle+ keys.

### planner

API for listing learning objects to display on the student planner and calendar

- **`canvas-cli planner index-other`** - Retrieve the paginated list of objects to be shown on the planner for the
current user with the associated planner override to override an item's
visibility if set.

Planner items for a student may also be retrieved by a linked observer. Use
the path that accepts a user_id and supply the student's id.
- **`canvas-cli planner overrides-create`** - Create a planner override for the current user
- **`canvas-cli planner overrides-destroy`** - Delete a planner override for the current user
- **`canvas-cli planner overrides-index`** - Retrieve a planner override for the current user
- **`canvas-cli planner overrides-show`** - Retrieve a planner override for the current user
- **`canvas-cli planner overrides-update`** - Update a planner override's visibilty for the current user

### planner-notes

API for creating, accessing and updating Planner Notes. PlannerNote are used
to set reminders and notes to self about courses or general events.

- **`canvas-cli planner-notes create`** - Create a planner note for the current user
- **`canvas-cli planner-notes destroy`** - Delete a planner note for the current user
- **`canvas-cli planner-notes index`** - Retrieve the paginated list of planner notes

Retrieve planner note for a user
- **`canvas-cli planner-notes show`** - Retrieve a planner note for the current user
- **`canvas-cli planner-notes update`** - Update a planner note for the current user

### progress

API for querying the progress of asynchronous API operations.

- **`canvas-cli progress show`** - Return completion and status information about an asynchronous job

### question-banks

Copyright (C) 2011 - present Instructure, Inc.

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

- **`canvas-cli question-banks assessment-index`** - Returns the paginated list of question banks for a given context.
- **`canvas-cli question-banks assessment-show`** - Returns the question bank with the given id

### rich-content

Manage rich content

- **`canvas-cli rich-content api-generate`** - Generates a rich content.

### rubrics

Manage rubrics

- **`canvas-cli rubrics api-upload-template`** - Returns a CSV template file that can be used to import rubrics into Canvas.

### sections

API for accessing section information.

- **`canvas-cli sections destroy`** - Delete an existing section.  Returns the former Section.
- **`canvas-cli sections show-other`** - Gets details about a specific section
- **`canvas-cli sections update`** - Modify an existing section.

### services

Manage services

- **`canvas-cli services api-show-kaltura-config`** - Return the config information for the Kaltura plugin in json format.
- **`canvas-cli services api-start-kaltura-session`** - Start a new Kaltura session, so that new media can be recorded and uploaded
to this Canvas instance's Kaltura instance.

### settings

Manage settings

- **`canvas-cli settings accounts-environment`** - Return a hash of global settings for the root account
This is the same information supplied to the web interface as +ENV.SETTINGS+.

### shared-brand-configs

This is how you can share Themes with other people in your account or
so you can come back to them later without having to apply them to your account

- **`canvas-cli shared-brand-configs destroy`** - Delete a SharedBrandConfig, which will unshare it so you nor anyone else in
your account will see it as an option to pick from.

### submissions

Manage submissions


### users

Manage users

- **`canvas-cli users communication-channels-delete-push-token`** - Delete a push notification endpoint
- **`canvas-cli users course-nicknames-clear`** - Remove all stored course nicknames.
- **`canvas-cli users course-nicknames-delete`** - Remove the nickname for the given course.
Subsequent course API calls will return the actual name for the course.
- **`canvas-cli users course-nicknames-index`** - Returns all course nicknames you have set.
- **`canvas-cli users course-nicknames-show`** - Returns the nickname for a specific course.
- **`canvas-cli users course-nicknames-update`** - Set a nickname for the given course. This will replace the course's name
in output of API calls you make subsequently, as well as in selected
places in the Canvas web user interface.
- **`canvas-cli users favorites-add-favorite-course`** - Add a course to the current user's favorites.  If the course is already
in the user's favorites, nothing happens. Canvas for Elementary subject
and homeroom courses can be added to favorites, but this has no effect in
the UI.
- **`canvas-cli users favorites-add-favorite-groups`** - Add a group to the current user's favorites.  If the group is already
in the user's favorites, nothing happens.
- **`canvas-cli users favorites-list-favorite-courses`** - Retrieve the paginated list of favorite courses for the current user. If the user has not chosen
any favorites, then a selection of currently enrolled courses will be returned.

See the {api:CoursesController#index List courses API} for details on accepted include[] parameters.
- **`canvas-cli users favorites-list-favorite-groups`** - Retrieve the paginated list of favorite groups for the current user. If the user has not chosen
any favorites, then a selection of groups that the user is a member of will be returned.
- **`canvas-cli users favorites-remove-favorite-course`** - Remove a course from the current user's favorites.
- **`canvas-cli users favorites-remove-favorite-groups`** - Remove a group from the current user's favorites.
- **`canvas-cli users favorites-reset-course-favorites`** - Reset the current user's course favorites to the default
automatically generated list of enrolled courses
- **`canvas-cli users favorites-reset-groups-favorites`** - Reset the current user's group favorites to the default
automatically generated list of enrolled group
- **`canvas-cli users groups-index`** - Returns a paginated list of active groups for the current user.
- **`canvas-cli users notification-preferences-update-all-for`** - Change the preferences for multiple notifications for a single communication channel at once
- **`canvas-cli users notification-preferences-update-all-for-2`** - Change the preferences for multiple notifications for a single communication channel at once
- **`canvas-cli users notification-preferences-update-for`** - Change the preference for a single notification for a single communication channel
- **`canvas-cli users notification-preferences-update-for-2`** - Change the preference for a single notification for a single communication channel
- **`canvas-cli users notification-preferences-update-preferences-by-category`** - Change the preferences for multiple notifications based on the category for a single communication channel
- **`canvas-cli users page-views-batch-query`** - Initiates an asynchronous query for page views data across multiple users.
This method enqueues a background job to process the batch page views query and returns
a polling URL that can be used to check the query status and retrieve results when ready.

As this is a beta endpoint, it is subject to change or removal at any time without the standard notice periods outlined in the API policy.
- **`canvas-cli users page-views-batch-query-results`** - Retrieves the results of a completed batch page views query. Returns the data in the
format specified when the query was initiated (CSV or JSON). The response may
be compressed with gzip encoding.

As this is a beta endpoint, it is subject to change or removal at any time without the standard notice periods outlined in the API policy.
- **`canvas-cli users page-views-poll-batch-query`** - Checks the status of a previously initiated batch page views query. Returns the current
processing status and provides a result URL when the query is complete.

As this is a beta endpoint, it is subject to change or removal at any time without the standard notice periods outlined in the API policy.
- **`canvas-cli users pseudonyms-forgot-password`** - Given a user email, generate a nonce and email it to the user
- **`canvas-cli users service-credentials-activity-stream-for`** - Returns the current user's global activity stream, paginated.

There are many types of objects that can be returned in the activity
stream. All object types have the same basic set of shared attributes:
  !!!javascript
  {
    'created_at': '2011-07-13T09:12:00Z',
    'updated_at': '2011-07-25T08:52:41Z',
    'id': 1234,
    'title': 'Stream Item Subject',
    'message': 'This is the body text of the activity stream item. It is plain-text, and can be multiple paragraphs.',
    'type': 'DiscussionTopic|Conversation|Message|Submission|Conference|Collaboration|AssessmentRequest...',
    'read_state': false,
    'context_type': 'course', // course|group
    'course_id': 1,
    'group_id': null,
    'html_url': "http://..." // URL to the Canvas web UI for this stream item
  }

In addition, each item type has its own set of attributes available.

DiscussionTopic:

  !!!javascript
  {
    'type': 'DiscussionTopic',
    'discussion_topic_id': 1234,
    'total_root_discussion_entries': 5,
    'require_initial_post': true,
    'user_has_posted': true,
    'root_discussion_entries': {
      ...
    }
  }

For DiscussionTopic, the message is truncated at 4kb.

Announcement:

  !!!javascript
  {
    'type': 'Announcement',
    'announcement_id': 1234,
    'total_root_discussion_entries': 5,
    'require_initial_post': true,
    'user_has_posted': null,
    'root_discussion_entries': {
      ...
    }
  }

For Announcement, the message is truncated at 4kb.

Conversation:

  !!!javascript
  {
    'type': 'Conversation',
    'conversation_id': 1234,
    'private': false,
    'participant_count': 3,
  }

Message:

  !!!javascript
  {
    'type': 'Message',
    'message_id': 1234,
    'notification_category': 'Assignment Graded'
  }

Submission:

Returns an {api:Submissions:Submission Submission} with its Course and Assignment data.

Conference:

  !!!javascript
  {
    'type': 'Conference',
    'web_conference_id': 1234
  }

Collaboration:

  !!!javascript
  {
    'type': 'Collaboration',
    'collaboration_id': 1234
  }

AssessmentRequest:

  !!!javascript
  {
    'type': 'AssessmentRequest',
    'assessment_request_id': 1234
  }
- **`canvas-cli users service-credentials-activity-stream-other`** - Returns the current user's global activity stream, paginated.

There are many types of objects that can be returned in the activity
stream. All object types have the same basic set of shared attributes:
  !!!javascript
  {
    'created_at': '2011-07-13T09:12:00Z',
    'updated_at': '2011-07-25T08:52:41Z',
    'id': 1234,
    'title': 'Stream Item Subject',
    'message': 'This is the body text of the activity stream item. It is plain-text, and can be multiple paragraphs.',
    'type': 'DiscussionTopic|Conversation|Message|Submission|Conference|Collaboration|AssessmentRequest...',
    'read_state': false,
    'context_type': 'course', // course|group
    'course_id': 1,
    'group_id': null,
    'html_url': "http://..." // URL to the Canvas web UI for this stream item
  }

In addition, each item type has its own set of attributes available.

DiscussionTopic:

  !!!javascript
  {
    'type': 'DiscussionTopic',
    'discussion_topic_id': 1234,
    'total_root_discussion_entries': 5,
    'require_initial_post': true,
    'user_has_posted': true,
    'root_discussion_entries': {
      ...
    }
  }

For DiscussionTopic, the message is truncated at 4kb.

Announcement:

  !!!javascript
  {
    'type': 'Announcement',
    'announcement_id': 1234,
    'total_root_discussion_entries': 5,
    'require_initial_post': true,
    'user_has_posted': null,
    'root_discussion_entries': {
      ...
    }
  }

For Announcement, the message is truncated at 4kb.

Conversation:

  !!!javascript
  {
    'type': 'Conversation',
    'conversation_id': 1234,
    'private': false,
    'participant_count': 3,
  }

Message:

  !!!javascript
  {
    'type': 'Message',
    'message_id': 1234,
    'notification_category': 'Assignment Graded'
  }

Submission:

Returns an {api:Submissions:Submission Submission} with its Course and Assignment data.

Conference:

  !!!javascript
  {
    'type': 'Conference',
    'web_conference_id': 1234
  }

Collaboration:

  !!!javascript
  {
    'type': 'Collaboration',
    'collaboration_id': 1234
  }

AssessmentRequest:

  !!!javascript
  {
    'type': 'AssessmentRequest',
    'assessment_request_id': 1234
  }
- **`canvas-cli users service-credentials-activity-stream-summary`** - Returns a summary of the current user's global activity stream.
- **`canvas-cli users service-credentials-api-show`** - Shows details for user.

Also includes an attribute "permissions", a non-comprehensive list of permissions for the user.
Example:
  !!!javascript
  "permissions": {
   "can_update_name": true, // Whether the user can update their name.
   "can_update_avatar": false, // Whether the user can update their avatar.
   "limit_parent_app_web_access": false // Whether the user can interact with Canvas web from the Canvas Parent app.
  }
- **`canvas-cli users service-credentials-expire-mobile-sessions-other`** - Permanently expires any active mobile sessions, forcing them to re-authorize.

The route that takes a user id will expire mobile sessions for that user.
The route that doesn't take a user id will expire mobile sessions for *all* users
in the institution (except for account administrators if +skip_admins+ is given).
- **`canvas-cli users service-credentials-ignore-all-stream-items`** - Hide all stream items for the user
- **`canvas-cli users service-credentials-ignore-stream-item`** - Hide the given stream item.
- **`canvas-cli users service-credentials-pandata-events-token`** - Returns a jwt auth and props token that can be used to send events to
Pandata.

NOTE: This is currently only available to the mobile developer keys.
- **`canvas-cli users service-credentials-todo-item-count`** - Counts of different todo items such as the number of assignments needing grading as well as the number of assignments needing submitting.

There is a limit to the number of todo items this endpoint will count.
It will only look at the first 100 todo items for the user. If the user has more than 100 todo items this count may not be reliable.
The largest reliable number for both counts is 100.
- **`canvas-cli users service-credentials-todo-items`** - A paginated list of the current user's list of todo items.

There is a limit to the number of items returned.

The `ignore` and `ignore_permanently` URLs can be used to update the user's
preferences on what items will be displayed.
Performing a DELETE request against the `ignore` URL will hide that item
from future todo item requests, until the item changes.
Performing a DELETE request against the `ignore_permanently` URL will hide
that item forever.
- **`canvas-cli users service-credentials-upcoming-events`** - A paginated list of the current user's upcoming events.
- **`canvas-cli users service-credentials-update`** - Modify an existing user. To modify a user's login, see the documentation for logins.


## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
canvas-cli account-calendars api-index

# JSON for scripting and agents
canvas-cli account-calendars api-index --json

# Filter to specific fields
canvas-cli account-calendars api-index --json --select id,name,status

# Dry run — show the request without sending
canvas-cli account-calendars api-index --dry-run

# Agent mode — JSON + compact + no prompts in one flag
canvas-cli account-calendars api-index --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
- **Filterable** - `--select id,name` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries and `--ignore-missing` to delete retries when a no-op success is acceptable
- **Confirmable** - `--yes` for explicit confirmation of destructive actions
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Offline-friendly** - sync/search commands can use the local SQLite store when available
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `4` auth error, `5` API error, `7` rate limited, `10` config error.

## Use with Claude Code

Install the focused skill — it auto-installs the CLI on first invocation:

```bash
npx skills add mvanhorn/printing-press-library/cli-skills/pp-canvas -g
```

Then invoke `/pp-canvas <query>` in Claude Code. The skill is the most efficient path — Claude Code drives the CLI directly without an MCP server in the middle.

<details>
<summary>Use as an MCP server in Claude Code (advanced)</summary>

If you'd rather register this CLI as an MCP server in Claude Code, install the MCP binary first:


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Then register it:

```bash
claude mcp add canvas canvas-mcp -e CANVAS_LMS_TOKEN=<your-token>
```

</details>

## Use with Claude Desktop

This CLI ships an [MCPB](https://github.com/modelcontextprotocol/mcpb) bundle — Claude Desktop's standard format for one-click MCP extension installs (no JSON config required).

To install:

1. Download the `.mcpb` for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/canvas-current).
2. Double-click the `.mcpb` file. Claude Desktop opens and walks you through the install.
3. Fill in `CANVAS_LMS_TOKEN` when Claude Desktop prompts you.

Requires Claude Desktop 1.0.0 or later. Pre-built bundles ship for macOS Apple Silicon (`darwin-arm64`) and Windows (`amd64`, `arm64`); for other platforms, use the manual config below.

<details>
<summary>Manual JSON config (advanced)</summary>

If you can't use the MCPB bundle (older Claude Desktop, unsupported platform), install the MCP binary and configure it manually.


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "canvas": {
      "command": "canvas-mcp",
      "env": {
        "CANVAS_LMS_TOKEN": "<your-key>"
      }
    }
  }
}
```

</details>

## Health Check

```bash
canvas-cli doctor
```

Verifies configuration, credentials, and connectivity to the API.

## Configuration

Config file: `~/.config/canvas-lms-pp-cli/config.toml`

Static request headers can be configured under `headers`; per-command header overrides take precedence.

Environment variables:

| Name | Kind | Required | Description |
| --- | --- | --- | --- |
| `CANVAS_LMS_TOKEN` | per_call | Yes | Set to your API credential. |

## Troubleshooting
**Authentication errors (exit code 4)**
- Run `canvas-cli doctor` to check credentials
- Verify the environment variable is set: `echo $CANVAS_LMS_TOKEN`
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

---

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press)
