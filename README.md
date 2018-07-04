# ses2case

Amazon SES can send you detailed notifications about your bounces, complaints, and deliveries.

This lambda function posts with an authorisation header these notification
messages to the Frontend (case), for inserting into the meteor/mongo User
collection that the user is invalid etc.

* https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notification-contents.html#bounce-types
* https://github.com/unee-t/frontend/issues/334
