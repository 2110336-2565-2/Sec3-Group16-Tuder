package datas

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/google/uuid"
)

func InsertIssueReport(client *ent.Client, ctx context.Context) []*ent.IssueReport {

	// Create users
	report1 := CreateIssueReport(client, "Test1", "Issue test data1", "No contact", time.Now())
	report2 := CreateIssueReport(client, "Test2", "Issue test data2", "081-234-5678", time.Now())
	report3 := CreateIssueReport(client, "Test3", "Issue test data3", "test@data.com", time.Now())

	reports, err := client.IssueReport.CreateBulk(report1, report2, report3).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating issue report: %v", err)
	}

	for _, report := range reports {
		fmt.Println("Issue report created: ", report.ID)
	}

	return reports
}

func CreateIssueReport(
	client *ent.Client,
	title string,
	description string,
	contact string,
	reportDate time.Time,
) *ent.IssueReportCreate {

	// Gen UUID and Hash Password
	issueReportId := uuid.New()

	issueReport := client.IssueReport.Create().
		SetID(issueReportId).
		SetTitle(title).
		SetDescription(description).
		SetContact(contact).
		SetReportDate(reportDate)

	return issueReport
}
