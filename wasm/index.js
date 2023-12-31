const goWasm = new Go();

WebAssembly.instantiateStreaming(
  fetch("xml2json.wasm"),
  goWasm.importObject
).then((result) => {
  goWasm.run(result.instance);
});

const type = "go";

const biggerXml = `<SaveContactRestRequest xmlns="http://schemas.datacontract.org/2004/07/SmashFly.WebServices.ContactManagerService.v2">
    <Contact>
        <Address1>1 Main Street</Address1>
        <Address2>Unit 2</Address2>
        <City>Boston</City>
        <Company>3COM</Company>
        <ContactListId>999999999</ContactListId>
        <Country>US</Country>
        <Education>
            <ContactEducation>
                <Degree>Bachelors</Degree>
                <GradYear>1988</GradYear>
                <HasGraduated>true</HasGraduated>
                <Major>Electrical Engineering</Major>
                <School>MIT</School>
            </ContactEducation>
            <ContactEducation>
                <Degree>Masters</Degree>
                <GradYear>1992</GradYear>
                <HasGraduated>true</HasGraduated>
                <Major>Business</Major>
                <School>Harvard</School>
            </ContactEducation>
        </Education>
        <EducationCurrent>BachelorofScience</EducationCurrent>
        <Email>gwashington7@invalidemail.com</Email>
        <EventListId>999999999</EventListId>
        <Experience>
            <ContactExperience>
                <Company>3COM</Company>
                <EndMonth>2</EndMonth>
                <EndYear>1998</EndYear>
                <JobTitle>Senior Engineer</JobTitle>
                <StartMonth>1</StartMonth>
                <StartYear>1992</StartYear>
                <State>MA</State>
                <Supervisor>Bob Smith</Supervisor>
                <SupervisorTitle>CTO</SupervisorTitle>
            </ContactExperience>
            <ContactExperience>
                <Company>EMC</Company>
                <EndMonth>7</EndMonth>
                <EndYear>2005</EndYear>
                <JobTitle>Business Analyst</JobTitle>
                <StartMonth>1</StartMonth>
                <StartYear>2003</StartYear>
            </ContactExperience>
        </Experience>
        <ExperienceCurrent>Experienced</ExperienceCurrent>
        <ExternalResumeId>9999999</ExternalResumeId>
        <FacebookProfile>https://www.facebook.com/smashflytechnologies</FacebookProfile>
        <FileToFolderPath>Specialized Sourcing/Veterans</FileToFolderPath>
        <FirstName>George</FirstName>
        <HomeEmail>gwashington@aol.com</HomeEmail>
        <IsEmployee>true</IsEmployee>
        <JobCode>19</JobCode>
        <JobId>LOS99710</JobId>
        <JobListStatus>Interviewing</JobListStatus>
        <JobListStatusNote>Great candidate</JobListStatusNote>
        <JobTitle>Senior Engineer</JobTitle>
        <LastName>Washington</LastName>
        <MobileNumber>9787931633</MobileNumber>
        <Notes>
            <ContactNote>
                <Note>George was a good president</Note>
                <Reminder>2013-08-31T11:20:00</Reminder>
            </ContactNote>
        </Notes>
        <Phone>978-342-3442</Phone>
        <PostalCode>01775</PostalCode>
        <ProfileURL>https://www.linkedin.com/company/846055</ProfileURL>
        <Resume>
           3200 Mount Vernon Hwy, Mt Vernon, VA SUMMARY: Former General
        </Resume>
        <ResumeBin>
            <Content xmlns="http://schemas.datacontract.org/2004/07/SmashFly.WebServices.ContactManagerService.v2.DataContracts">
              Some Content
            </Content>
            <Name xmlns="http://schemas.datacontract.org/2004/07/SmashFly.WebServices.ContactManagerService.v2.DataContracts">Resume.txt</Name>
        </ResumeBin>
        <SecondarySource>Indeed</SecondarySource>
        <State>MA</State>
        <TDSearchFolder>Text goes here</TDSearchFolder>
        <Tags>
            <ContactTag>
                <Access>Public</Access>
                <Tag>Engineering</Tag>
            </ContactTag>
            <ContactTag>
                <Access>Public</Access>
                <Tag>Java</Tag>
            </ContactTag>
        </Tags>
        <TwitterProfile>https://twitter.com/smashfly</TwitterProfile>
        <UDF>
            <ContactUDFData>
                <FieldName>ShortTextField4</FieldName>
                <Values>
                    <string xmlns="http://schemas.microsoft.com/2003/10/Serialization/Arrays">CivilEngineering</string>
                    <string xmlns="http://schemas.microsoft.com/2003/10/Serialization/Arrays">ElectricalEngineering</string>
                </Values>
            </ContactUDFData>
            <ContactUDFData>
                <FieldName>ShortTextField7</FieldName>
                <Values>
                    <string xmlns="http://schemas.microsoft.com/2003/10/Serialization/Arrays">Construction</string>
                    <string xmlns="http://schemas.microsoft.com/2003/10/Serialization/Arrays">Facilities_Mgmt</string>
                </Values>
            </ContactUDFData>
        </UDF>
        <UserScore>2.56743233E+15</UserScore>
        <WebURL>http://www.smashfly.com</WebURL>
        <WorkPhone>781-222-4422</WorkPhone>
        <_CanReceiveEmailImpl>true</_CanReceiveEmailImpl>
        <_CanReceiveSMSImpl>true</_CanReceiveSMSImpl>
        <_IsApplicantImpl>true</_IsApplicantImpl>
        <CanReceiveEmail>true</CanReceiveEmail>
        <CanReceiveSMS>true</CanReceiveSMS>
        <IsApplicant>true</IsApplicant>
    </Contact>
    <ContactDBId>99999999</ContactDBId>
    <Password>YourPassword</Password>
    <UserName>YourUsername</UserName>
</SaveContactRestRequest>`;

function ParseXml(xml) {
  if (type === "go") {
    const json = XmlToJson(xml);
    return JSON.parse(json);
  } else {
    return xml2json(xml);
  }
}

function RunTest() {
  const start = new Date().getTime();
  for (let x = 0; x < 100; x++) {
    if (x % 10 === 0) {
      // console.log(x);
    }
    const json = ParseXml(biggerXml);
  }
  const end = new Date().getTime();

  console.log("Result: ", end - start + "ms");
}
