<?php
if (isset($_REQUEST['formtype']) && $_REQUEST['formtype'] == 'form') {
	if (isset($_REQUEST['formname'])) {
		
	}
} elseif (isset($_REQUEST['formtype']) && $_REQUEST['formtype'] == 'widget') {
	if (isset($_REQUEST['formname'])) {
		$formname = $_REQUEST['formname'];
		$field_003 = $_REQUEST['field_003'];
		$field_004 = $_REQUEST['field_004'];
		
		$mailAddress = 'yourmail@mail.com';
		$subject = 'Subject';
		$msg = "Message $field_003\r\n $field_004";
		$headers = "MIME-Version: 1.0\r\n Content-type: text/plain; charset=utf-8\r\n From: " . $_REQUEST[$result->slug] . "\r\n Reply-To: " . $_REQUEST[$result->slug] . "\r\n X-Mailer: PHP/" . phpversion();
		
		mail($mailAddress, $subject, $msg, $headers);
	}
} elseif (isset($_REQUEST['formtype']) && $_REQUEST['formtype'] == 'contactfmain') {
	if (isset($_REQUEST['formname'])) {
		$formname = $_REQUEST['formname'];
		$contact_name = $_REQUEST['contact_name'];
		$contact_email = $_REQUEST['contact_email'];
		$contact_url = $_REQUEST['contact_url'];
		$contact_subject = $_REQUEST['contact_subject'];
		$contact_message = $_REQUEST['contact_message'];
		
		$mailAddress = 'yourmail@mail.com';
		$subject = 'Subject';
		$msg = "$contact_name\r\n $contact_email\r\n $contact_url\r\n $contact_subject\r\n $contact_message";
		$headers = "MIME-Version: 1.0\r\n Content-type: text/plain; charset=utf-8\r\n From: " . $_REQUEST[$result->slug] . "\r\n Reply-To: " . $_REQUEST[$result->slug] . "\r\n X-Mailer: PHP/" . phpversion();
		
		mail($mailAddress, $subject, $msg, $headers);
	}
} elseif (isset($_REQUEST['formtype']) && $_REQUEST['formtype'] == 'contactf_1') {
	if (isset($_REQUEST['formname'])) {
		$formname = $_REQUEST['formname'];
		$contact_email_1 = $_REQUEST['contact_email_1'];
		$contact_message_1 = $_REQUEST['contact_message_1'];
		
		$mailAddress = 'yourmail@mail.com';
		$subject = 'Subject';
		$msg = "$contact_email_1\r\n $contact_message_1";
		$headers = "MIME-Version: 1.0\r\n Content-type: text/plain; charset=utf-8\r\n From: " . $_REQUEST[$result->slug] . "\r\n Reply-To: " . $_REQUEST[$result->slug] . "\r\n X-Mailer: PHP/" . phpversion();
		
		mail($mailAddress, $subject, $msg, $headers);
	}
} elseif (isset($_REQUEST['formtype']) && $_REQUEST['formtype'] == 'contactf_2') {
	if (isset($_REQUEST['formname'])) {
		$formname = $_REQUEST['formname'];
		$contact_name_2 = $_REQUEST['contact_name_2'];
		$age_group = $_REQUEST['age_group'];
		$gender = $_REQUEST['gender'];
		$cmsms_which_statements_are_true_for1 = $_REQUEST['cmsms_which_statements_are_true_for1'];
		$cmsms_which_statements_are_true_for2 = $_REQUEST['cmsms_which_statements_are_true_for2'];
		$cmsms_which_statements_are_true_for3 = $_REQUEST['cmsms_which_statements_are_true_for3'];
		$cmsms_which_statements_are_true_for4 = $_REQUEST['cmsms_which_statements_are_true_for4'];
		$cmsms_which_statements_are_true_for5 = $_REQUEST['cmsms_which_statements_are_true_for5'];
		$contact_email = $_REQUEST['contact_email'];
		$contact_phone = $_REQUEST['contact_phone'];
		
		$mailAddress = 'yourmail@mail.com';
		$subject = 'Subject';
		$msg = "$contact_name_2\r\n $age_group\r\n $gender\r\n $cmsms_which_statements_are_true_for1\r\n $cmsms_which_statements_are_true_for2\r\n $cmsms_which_statements_are_true_for3\r\n $cmsms_which_statements_are_true_for4\r\n $cmsms_which_statements_are_true_for5\r\n $contact_email\r\n $contact_phone";
		$headers = "MIME-Version: 1.0\r\n Content-type: text/plain; charset=utf-8\r\n From: " . $_REQUEST[$result->slug] . "\r\n Reply-To: " . $_REQUEST[$result->slug] . "\r\n X-Mailer: PHP/" . phpversion();
		
		mail($mailAddress, $subject, $msg, $headers);
	}
} elseif (isset($_REQUEST['formtype']) && $_REQUEST['formtype'] == 'contactfmain_2') {
	if (isset($_REQUEST['formname'])) {
		$formname = $_REQUEST['formname'];
		$contact_name = $_REQUEST['contact_name'];
		$contact_email = $_REQUEST['contact_email'];
		$contact_subject = $_REQUEST['contact_subject'];
		$contact_message = $_REQUEST['contact_message'];
		
		$mailAddress = 'yourmail@mail.com';
		$subject = 'Subject';
		$msg = "$contact_name\r\n $contact_email\r\n $contact_subject\r\n $contact_message";
		$headers = "MIME-Version: 1.0\r\n Content-type: text/plain; charset=utf-8\r\n From: " . $_REQUEST[$result->slug] . "\r\n Reply-To: " . $_REQUEST[$result->slug] . "\r\n X-Mailer: PHP/" . phpversion();
		
		mail($mailAddress, $subject, $msg, $headers);
	}
}
?>